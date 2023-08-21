package user

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"

	"auth/internal/db"
	"auth/internal/models"
	"auth/internal/redisclient"
	"auth/pkg/contracts"
)

func RegisterUser(data contracts.RegisterUserRequest) error {
	// return "Success"
	user := models.User{
		ID:                 generateUUID(),
		TenantId:           data.ClientID,
		FirstName:          data.FirstName,
		LastName:           data.LastName,
		EmailID:            data.Email,
		MobileNumber:       data.Mobile,
		Country:            "US",
		NotificationStatus: true,
	}

	creds := models.Credientials{
		UserId:   user.ID,
		Password: data.Password,
	}

	var existingUser models.User
	result := db.DB.Where("email_id = ?", user.EmailID).First(&existingUser)
	if result.Error == nil {
		return errors.New("user with the same email already exists")
	}

	tx := db.DB.Begin() // Start a new transaction

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		return err
	}

	if err := tx.Create(&creds).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		return err
	}

	tx.Commit() // Commit the transaction if all operations are successful
	return nil
}

func ChangePassword(data contracts.ChangePassword) error {

	var user models.User
	result := db.DB.Where("email_id = ? and tenant_id = ?", data.Email, data.ClientID).First(&user)
	if result.Error != nil {
		return errors.New("user with the email and tenant does not exist")
	}

	var cred models.Credientials
	result = db.DB.Where("user_id = ?", user.ID).First(&cred)
	if result.Error != nil {
		return errors.New("unable to find the creds of the user")
	}

	// TODO: Addedd the hashing logic here.
	cred.Password = data.Password
	result = db.DB.Save(cred)
	if result.Error != nil {
		return errors.New("unable to update credentials")
	}
	return nil
}

func AddPermission(userId string, data contracts.UserPermissionRequest) error {
	var user models.User
	result := db.DB.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return errors.New("unable to find the user")
	}
	var resourcePermission models.ResourcePermission
	result = db.DB.Where("permission_name = ? and permission_action = ?", data.PermissionName, data.PermissionAction).First(&resourcePermission)
	if result.Error != nil {
		return errors.New("unable to find the user")
	}

	userPermission := models.UserPermissions{
		PermissionId: resourcePermission.ID,
		UserId:       user.ID,
	}
	db.DB.Create(&userPermission)

	// Recreate the permission against the user who is currently logged in
	ctx := context.Background()
	token, err := redisclient.RedisClient.Get(ctx, user.ID).Result()
	if err == nil {
		CacheUserPermission(token, user.ID)
	}

	return nil
}

func RevokePermission(userId string, data contracts.UserPermissionRequest) error {
	var user models.User
	result := db.DB.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return errors.New("unable to find the user")
	}
	var resourcePermission models.ResourcePermission
	result = db.DB.Where("permission_name = ? and permission_action = ?", data.PermissionName, data.PermissionAction).First(&resourcePermission)
	if result.Error != nil {
		return errors.New("unable to find the user")
	}

	var userPermission models.UserPermissions
	result = db.DB.Where("permission_id = ? and user_id = ?", resourcePermission.ID, user.ID).First(&userPermission)
	if result.Error != nil {
		return errors.New("unable to get user permission")
	}
	db.DB.Delete(&userPermission)

	// Recreate the permission against the user who is currently logged in
	ctx := context.Background()
	token, err := redisclient.RedisClient.Get(ctx, user.ID).Result()
	if err == nil {
		CacheUserPermission(token, user.ID)
	}
	return nil
}

func CacheUserPermission(token, userId string) error {
	var userPermissions []models.UserPermissions
	result := db.DB.Preload("User").Preload("ResourcePermission").Preload("ResourcePermission.ResourceGroup").Where("user_id = ?", userId).Find(&userPermissions)
	if result.Error != nil {
		return result.Error
	}
	permissionMap := make(map[string]map[string]string)
	isAdmin := false
	for _, userPermission := range userPermissions {
		if userPermission.User.IsAdmin {
			isAdmin = true
			break
		}
		name := userPermission.ResourcePermission.PermissionName
		action := userPermission.ResourcePermission.PermissionAction
		if permissionMap[userPermission.ResourcePermission.ResourceGroup.GroupName] == nil {
			permissionMap[userPermission.ResourcePermission.ResourceGroup.GroupName] = make(map[string]string)
		}
		permissionMap[userPermission.ResourcePermission.ResourceGroup.GroupName][action] = name
	}
	if isAdmin {
		permissionMap["Admin"] = make(map[string]string)
	}

	permissions, err := json.Marshal(permissionMap)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = redisclient.RedisClient.Set(ctx, token, string(permissions), 10*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func generateUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}

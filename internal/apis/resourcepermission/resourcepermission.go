package resourcepermission

import (
	"auth/internal/db"
	"auth/internal/models"

	"github.com/google/uuid"
)

func InsertResourcePermission(permissionName, permissionAction, groupName string) error {
	var resourceGroup models.ResourceGroup
	result := db.DB.Where("group_name = ?", groupName).First(&resourceGroup)
	if result.Error != nil {
		return result.Error
	}

	newPermission := models.ResourcePermission{
		PermissionName:   permissionName,
		PermissionAction: permissionAction,
		ResourceGroupID:  resourceGroup.Id,
	}

	result = db.DB.Create(&newPermission)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func generateUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}

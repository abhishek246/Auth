package resourcegroup

import (
	"auth/internal/db"
	"auth/internal/models"
	"errors"

	"github.com/google/uuid"
)

func InsertResourceGroup(groupName, groupDescription string) error {
	newGroup := models.ResourceGroup{
		Id:               generateUUID(),
		GroupName:        groupName,
		GroupDescription: groupDescription,
	}

	var group models.ResourceGroup
	check := db.DB.Where("group_name = ?", groupName).First(&group)
	if check.Error == nil {
		return errors.New("Group with name already exists")
	}

	result := db.DB.Create(&newGroup)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateResourceGroup(id, groupName, groupDescription string) error {
	var group models.ResourceGroup

	result := db.DB.Model(&group).Where("id = ?", id).Updates(models.ResourceGroup{
		GroupName:        groupName,
		GroupDescription: groupDescription,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteResourceGroup(id string) error {
	var group models.ResourceGroup

	result := db.DB.Where("id = ?", id).Delete(&group)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetResourceGroupByID(id string) (*models.ResourceGroup, error) {
	var group models.ResourceGroup

	result := db.DB.Where("id = ?", id).First(&group)
	if result.Error != nil {
		return nil, result.Error
	}

	return &group, nil
}

func generateUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}

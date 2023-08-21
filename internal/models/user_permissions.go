package models

import "gorm.io/gorm"

type UserPermissions struct {
	gorm.Model
	PermissionId       uint
	UserId             string
	ResourcePermission ResourcePermission `gorm:"foreignKey: PermissionId"`
	User               User               `gorm:"foreignKey: UserId"`
}

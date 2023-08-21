package models

import "gorm.io/gorm"

type ResourcePermission struct {
	gorm.Model
	PermissionName        string
	PermissionDescription string
	PermissionAction      string
	ResourceGroupID       string
	ResourceGroup         ResourceGroup `gorm:"foreignKey:ResourceGroupID"`
}

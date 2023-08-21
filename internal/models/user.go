package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 string
	TenantId           string
	FirstName          string
	LastName           string
	EmailID            string
	MobileNumber       string
	Country            string
	NotificationStatus bool
	IsAdmin            bool
}

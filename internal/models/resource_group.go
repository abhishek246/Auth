package models

import "gorm.io/gorm"

type ResourceGroup struct {
	gorm.Model
	Id               string
	GroupName        string
	GroupDescription string
}

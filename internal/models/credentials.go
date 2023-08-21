package models

import (
	"gorm.io/gorm"
)

type Credientials struct {
	gorm.Model
	UserId   string
	Password string
}

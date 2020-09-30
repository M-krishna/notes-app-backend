package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255" json:"first_name"`
	LastName  string `gorm:"size:255" json:"last_name"`
}

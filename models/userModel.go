package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required,min=8"`
}

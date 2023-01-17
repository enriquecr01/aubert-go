package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uint
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string
	Token    string
}

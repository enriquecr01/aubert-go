package models

import (
	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model

	ID       uint
	Name     string
	URL      string
	Username string
	Secret   string
	Notes    string
	Color    string
	UserId   int
}

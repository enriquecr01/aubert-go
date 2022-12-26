package models

import (
	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model

	Id       uint
	Name     string
	URL      string
	Username string
	Secret   string
	Notes    string
	Color    string
}

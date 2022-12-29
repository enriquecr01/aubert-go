package models

import (
	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model

	ID       uint   
	Name     string `json:"name"`
	URL      string `json:"url"`
	Username string `json:"user"`
	Secret   string `json:"secret"`
	Notes    string `json:"notes"`
	Color    string `json:"color"`
	UserId   int    `json:"user_id"`
}

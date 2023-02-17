package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model

	ID     uint
	Title  string `json:"title"`
	Note   string `json:"note"`
	Color  string `json:"color"`
	UserId int    `json:"user_id"`
}

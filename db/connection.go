package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var host = "162.241.60.131:3306"
var username = "thunderk_admin"
var password = "roaming32"
var database = "thunderk_aubert"

func DBConnection() {
	var error error
	DSN := `thunderk_admin:roaming32@tcp(162.241.60.131:3306)/thunderk_aubert?charset=utf8mb4&parseTime=True&loc=Local`
	DB, error = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}

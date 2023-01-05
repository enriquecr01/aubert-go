package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/enriquecr01/aubert-go/db"
	"github.com/enriquecr01/aubert-go/models"
)

func AddUser(c *gin.Context) {

	name := c.PostForm("name")
	email := c.PostForm("email")
	pass := c.PostForm("pass")
	message := "Registered Succesfully"
	status := 0

	newSecret := models.User{Name: name, Password: pass, Email: email}
	result := db.DB.Omit("updated_at", "deleted_at").Create(&newSecret)

	if result.Error != nil {
		status = 1
		message = "Error during insert"
	}

	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}

func ModifyUser(c *gin.Context) {

	name := c.PostForm("name")
	email := c.PostForm("email")
	pass := c.PostForm("pass")
	userId, err := strconv.Atoi(c.PostForm("userId"))
	message := "Updated Succesfully"
	status := 0

	if err != nil {
		message = "Error during parsing"
		status = 1
	} else {
		updatedUser := db.DB.Where("id = ?", userId).Updates(models.User{Name: name, Password: pass, Email: email})

		fmt.Printf("%+v\n", updatedUser)
	}

	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}

func Login(c *gin.Context) {

	email := c.PostForm("email")
	pass := c.PostForm("pass")
	status := 0
	message := ""
	var user models.User

	where := "email = '" + email + "' AND password = '" + pass + "' "
	db.DB.Select("id", "name", "email").Omit("CreatedAt", "UpdatedAt", "DeletedAt", "password").Where(where).Find(&user)

	if user.ID == 0 {
		status = 1
		message = "Login failed, review your credentials"
		c.JSON(401, gin.H{
			"status":  status,
			"message": message,
		})
	} else {
		c.JSON(200, gin.H{
			"status": status,
			"user":   user,
		})
	}
}

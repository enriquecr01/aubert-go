package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/enriquecr01/aubert-go/db"
	"github.com/enriquecr01/aubert-go/models"
)

func GetAllPasswords(c *gin.Context) {
	var secrets []models.Secret

	userId := c.Param("userid")
	where := "user_id = " + userId
	db.DB.Select("id", "name", "url", "username", "secret", "notes", "color").Omit("CreatedAt", "UpdatedAt", "DeletedAt").Where(where).Find(&secrets)

	c.JSON(200, gin.H{
		"secrets": &secrets,
	})
}

func SearchPasswords(c *gin.Context) {
	var secrets []models.Secret

	userId := c.Param("userid")
	term := c.Param("term")

	where := "user_id = " + userId + " AND name LIKE '%" + term + "%' "
	db.DB.Where(where).Find(&secrets)

	c.JSON(200, gin.H{
		"secrets": &secrets,
	})
}

func AddPassword(c *gin.Context) {

	name := c.PostForm("name")
	url := c.PostForm("url")
	username := c.PostForm("username")
	secret := c.PostForm("secret")
	notes := c.PostForm("notes")
	color := c.PostForm("color")
	userId, err := strconv.Atoi(c.PostForm("userId"))
	message := "Registered Succesfully"

	if err != nil {
		message = "Error during parsing"
	} else {
		newSecret := models.Secret{Name: name, URL: url, Username: username, Secret: secret, Notes: notes, Color: color, UserId: userId}
		result := db.DB.Omit("updated_at", "deleted_at").Create(&newSecret)

		if result.Error != nil {
			message = "Error during insert"
		}
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}

func ModifyPassword(c *gin.Context) {

	name := c.PostForm("name")
	url := c.PostForm("url")
	username := c.PostForm("username")
	secret := c.PostForm("secret")
	notes := c.PostForm("notes")
	color := c.PostForm("color")
	passId, err := strconv.Atoi(c.PostForm("passId"))
	message := "Updated Succesfully"

	if err != nil {
		message = "Error during parsing"
	} else {
		updatedSecret := db.DB.Where("id = ?", passId).Updates(models.Secret{Name: name, URL: url, Username: username, Notes: notes, Secret: secret, Color: color})

		fmt.Printf("%+v\n", updatedSecret)
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}

func DeletePassword(c *gin.Context) {

	passId := c.Param("passId")
	message := "Deleted Succesfully"

	deletedSecret := db.DB.Delete(&models.Secret{}, passId)

	if deletedSecret.Error != nil {
		message = "Error deleting"
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}

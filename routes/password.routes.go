package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/enriquecr01/aubert-go/db"
	"github.com/enriquecr01/aubert-go/models"
)

func SearchPasswords(c *gin.Context) {
	var secrets []models.Secret
	db.DB.Find(&secrets)

	c.JSON(200, gin.H{
		"message": &secrets,
	})
}

func AddPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add le mot de passe",
	})
}

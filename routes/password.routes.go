package routes

import (
	"github.com/gin-gonic/gin"
)

func SearchPasswords(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Te voy a dar un password!",
	})
}

func AddPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add le mot de passe",
	})
}

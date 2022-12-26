package routes

import (
	"github.com/gin-gonic/gin"
)

func HomeHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello me llamo Aubert!",
	})
}

func ExampleHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Salut ç'est un exemple",
	})
}

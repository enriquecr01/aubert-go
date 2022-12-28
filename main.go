package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enriquecr01/aubert-go/db"
	"github.com/enriquecr01/aubert-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Server initialized")
	StartGin()
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	db.DBConnection()

	router := gin.New()
	router.GET("/", routes.HomeHandle)
	router.GET("/example", routes.ExampleHandle)
	router.GET("/password/search/:userid", routes.GetAllPasswords)
	router.GET("/password/search/:userid/:term", routes.SearchPasswords)
	router.POST("/password/add", routes.AddPassword)
	router.PUT("/password/modify", routes.ModifyPassword)
	router.DELETE("/password/delete/:passId", routes.DeletePassword)

	port := os.Getenv("PORT")
	fmt.Println("Port", os.Getenv("PORT"))
	if port == "" {
		port = "8080"
		fmt.Println("Running in port", 8080)
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

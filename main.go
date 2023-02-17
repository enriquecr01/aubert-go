package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enriquecr01/aubert-go/db"
	"github.com/enriquecr01/aubert-go/initializers"
	"github.com/enriquecr01/aubert-go/middleware"
	"github.com/enriquecr01/aubert-go/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Server initialized")
	initializers.LoadEnvVariables()
	StartGin()
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	db.DBConnection()

	router := gin.New()

	router.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"*"},
		AllowMethods:    []string{"PUT", "POST", "GET", "OPTIONS", "DELETE"},
		AllowHeaders:    []string{"Origin", "Authorization"},
		AllowAllOrigins: true,
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/", routes.HomeHandle)
	router.GET("/example", routes.ExampleHandle)
	// Password routes
	router.GET("/password/search/:userid", middleware.RequireAuth, routes.GetAllPasswords)
	router.GET("/password/search/:userid/:term", middleware.RequireAuth, routes.SearchPasswords)
	router.POST("/password/add", middleware.RequireAuth, routes.AddPassword)
	router.PUT("/password/modify", middleware.RequireAuth, routes.ModifyPassword)
	router.DELETE("/password/delete/:passId", middleware.RequireAuth, routes.DeletePassword)
	// User routes
	router.POST("/user/add", routes.AddUser)
	router.PUT("/user/modify", routes.ModifyUser)
	router.POST("/login", routes.Login)
	// Note routes
	router.GET("/note/search/:userid", middleware.RequireAuth, routes.GetAllNotes)
	router.GET("/note/search/:userid/:term", middleware.RequireAuth, routes.SearchNotes)
	router.POST("/note/add", middleware.RequireAuth, routes.AddNote)
	router.PUT("/note/modify", middleware.RequireAuth, routes.ModifyNote)
	router.DELETE("/note/delete/:noteId", middleware.RequireAuth, routes.DeleteNote)

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

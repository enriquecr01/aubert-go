package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello me llamo Aubert!"))
}

func exampleHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Salut ç'est un exemple"))
}

func main() {
	// fmt.Println("Server initialized")
	StartGin()

	// r := mux.NewRouter()

	// r.HandleFunc("/", homeHandle)
	// r.HandleFunc("/example", exampleHandle)

	// http.ListenAndServe(":80", r)

}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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

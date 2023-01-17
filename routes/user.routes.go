package routes

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/enriquecr01/aubert-go/db"
	"github.com/enriquecr01/aubert-go/models"
	"github.com/golang-jwt/jwt/v4"
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
	db.DB.Select("id", "name", "email", "password").Omit("CreatedAt", "UpdatedAt", "DeletedAt").Where(where).Find(&user)

	// claims := jwt.MapClaims{
	// 	"sub": user.ID,
	// 	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	// }

	// claims := {
	// 	"id":    user.ID,
	// 	"email": user.Email,
	// 	"expires": {
	// 		"ExpiresAt": time.Now().Add(time.Hour * 48).Unix(),
	// 		"IssuedAt":  time.Now().Unix(),
	// 	},
	// }

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["user"] = user.ID

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//token := jwt.NewWithClaims(jwt.SigningMethodHMAC, jwt.MapClaims{
	//	"sub": user.ID,
	//	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	//})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	db.DB.Where("id = ?", user.ID).Updates(models.User{Token: tokenString})

	if user.ID == 0 {
		status = 1
		message = "Login failed, review your credentials"
		c.JSON(401, gin.H{
			"status":  status,
			"message": message,
		})
	} else if err != nil {
		status = 1
		message = "Login failed creating token"
		fmt.Println("Error creating token ", err)
		c.JSON(401, gin.H{
			"status":  status,
			"message": message,
		})
	} else {
		c.JSON(200, gin.H{
			"status": status,
			"user":   user,
			"token":  tokenString,
		})
	}
}

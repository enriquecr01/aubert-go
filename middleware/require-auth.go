package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/enriquecr01/aubert-go/db"
	"github.com/enriquecr01/aubert-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {

	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	var user models.User

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signign Method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":       1,
			"message":      "Error decoding token",
			"errorMessage": err,
		})
	}

	// getting the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprint(claims["user"])

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "Error getting claims",
		})
	}

	where := "id = " + userId
	searchUser := db.DB.Select("id", "name", "email", "token").Omit("CreatedAt", "UpdatedAt", "DeletedAt", "password").Where(where).Find(&user)

	fmt.Println("XDDDD", searchUser)

	if user.ID != 0 {

		if tokenString == user.Token {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  1,
				"message": "Error the token it's not the same",
			})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "Error getting user",
		})
	}

	//c.Next()
}

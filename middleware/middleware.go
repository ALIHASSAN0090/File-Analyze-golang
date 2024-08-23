package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key")

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println("Token String:", tokenString) // Debug statement

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			fmt.Println("Error parsing token:", err) // Debug statement
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Println("Token is not valid") // Debug statement
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Error asserting claims") // Debug statement
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid or expired token"})
			c.Abort()
			return
		}

		fmt.Println("Token Claims:", claims) // Debug statement
		c.Set("userName", claims["name"])

		c.Next()
	}
}

func GenerateToken(userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["name"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println("Error generating token:", err) // Debug statement
		return "", err
	}

	fmt.Println("Generated Token:", tokenString) // Debug statement
	return tokenString, nil
}

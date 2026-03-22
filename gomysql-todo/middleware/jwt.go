package middleware

import (
	"fmt"
	"gomysql-todo/models"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{"error": "token missing"})
			c.Abort()
			return
		}

		JWT_SECRET := os.Getenv("JWT_SECRET")
		if JWT_SECRET == "" {
			c.JSON(401, gin.H{"message": "JWT key not found"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		secret := []byte(JWT_SECRET)
		claim := &models.Claim{}

		fmt.Println(secret)
		fmt.Println(tokenStr)
		token, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claim.Username)
		c.Set("user_id", uint(claim.UserID))
		c.Next()
	}
}

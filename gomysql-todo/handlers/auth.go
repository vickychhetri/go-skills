package handlers

import (
	"errors"
	"gomysql-todo/config"
	"gomysql-todo/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": "wrong input"})
		return
	}

	var user models.User

	result := config.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(400, gin.H{"message": "user not exist"})
			return
		} else {
			c.JSON(400, gin.H{"message": result.Error})
			return
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid credentials"})
		return
	}

	//generate token and return token
	exp := time.Now().Add(time.Hour * 2)
	claim := models.Claim{
		Username: user.Username,
		UserID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	JWT_SECRET := os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		c.JSON(400, gin.H{"message": "JWT key not found"})
		return
	}

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"token": tokenString})
}

func Register(c *gin.Context) {

	type Request struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	var req Request
	var user models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": "wrong data"})
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(400, gin.H{"message": "Password and Confirm Password are not same"})
		return
	}

	user.Username = req.Username
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 0)
	if err != nil {
		c.JSON(400, gin.H{"message": "Password Encryption issue"})
		return
	}
	user.Password = string(hash)
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{"message": "Unable to create user"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(400, gin.H{"message": "No user created"})
		return
	}

	c.JSON(200, gin.H{
		"message": "user created successfully",
	})
}

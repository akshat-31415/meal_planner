package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"mealplan/database"
	"mealplan/models"
	"mealplan/utils"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Register(c *gin.Context){
	var Input struct{
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&Input); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(Input.Password)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"could not hash password"})
		return
	}

	user := models.User{
		Name: Input.Name,
		Email: Input.Email,
		Password: hashedPassword,
	}

	if err:=database.DB.Create(&user).Error; err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message": "User registered succeessfully"})
}

func Login(c *gin.Context){
	var Input struct{
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&Input); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err:=database.DB.Where("email=?",Input.Email).First(&user).Error;err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(Input.Password,user.Password){
		c.JSON(http.StatusUnauthorized,gin.H{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(10*time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Could not create token"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"token": tokenString})
}

func Profile(c *gin.Context){
	userID, exists := c.Get("user")
	if !exists{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized"})
		return
	}

	var user models.User
    if err := database.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Don't send password in response
    user.Password = ""

	c.JSON(http.StatusOK,gin.H{"user":user})
}
package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnxvi/nyalist/models"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	DB *gorm.DB
}

func NewController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}

}

func (ac *AuthController) Login(c *gin.Context) {
	var loginData LoginType
	var user models.User
	if err := c.BindJSON(&loginData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ac.DB.First(&user, "email = ?", loginData.Email)
	if user.Email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot login with the provided credentials"},
		)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(loginData.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot login with the provided credentials"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": user.AuthToken})
}

func (ac *AuthController) Signup(c *gin.Context) {
	var user UserType
	var tokenManager TokenManager
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := user.validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	fmt.Println(hashed_password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ac.DB.Create(&models.User{Username: user.Username,
		Email: user.Email, Password: string(hashed_password),
		AuthToken: tokenManager.generate_token(),
	})
	c.JSON(http.StatusOK, gin.H{
		"message":  "User Created",
		"username": user.Username,
	})
}

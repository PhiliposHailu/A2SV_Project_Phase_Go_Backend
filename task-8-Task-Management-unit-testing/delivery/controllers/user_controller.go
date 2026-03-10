package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipos/api/domain"
)
type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(us domain.UserUsecase) *UserController {
	return &UserController{
		userUsecase: us,
	}
}


func (uc *UserController) Register(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	fmt.Println(user)
	err := uc.userUsecase.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	token, err := uc.userUsecase.Login(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
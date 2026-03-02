package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipos/api/data"
	"github.com/philipos/api/models"
)

func Register(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := data.RegisterService(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser.Password = "*******"
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    newUser,
	})


}

func Login(c * gin.Context) {
	var info models.User
	if err := c.BindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	token, err := data.LoginService(info.Username, info.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
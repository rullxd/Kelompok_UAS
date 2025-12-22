package controllers

import (
	"github.com/gin-gonic/gin"
	"pengaduan-kampus/services"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.ShouldBindJSON(&input)

	token, err := services.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

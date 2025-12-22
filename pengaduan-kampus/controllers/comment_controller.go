package controllers

import (
	"github.com/gin-gonic/gin"
	"pengaduan-kampus/config"
	"pengaduan-kampus/models"
	"pengaduan-kampus/utils"
)

func AddComment(c *gin.Context) {
	var input struct {
		Message string `json:"message"`
	}
	c.ShouldBindJSON(&input)

	config.DB.Create(&models.Comment{
		ComplaintID: utils.ParseUint(c.Param("id")),
		UserID:      c.GetUint("user_id"),
		Message:     input.Message,
	})

	c.JSON(200, gin.H{"message": "Komentar ditambahkan"})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"pengaduan-kampus/config"
	"pengaduan-kampus/models"
)

func ActionHistory(c *gin.Context) {
	var logs []models.ActionLog
	id := c.Param("id")

	config.DB.
		Where("complaint_id = ?", id).
		Order("created_at desc").
		Find(&logs)

	c.JSON(200, logs)
}


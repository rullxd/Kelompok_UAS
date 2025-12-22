package controllers

import (
	"github.com/gin-gonic/gin"
	"pengaduan-kampus/services"
)

func AllComplaints(c *gin.Context) {
	c.JSON(200, services.GetAllComplaints())
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"pengaduan-kampus/services"
)

func ComplaintDetail(c *gin.Context) {
	data := services.GetComplaintDetail(c.Param("id"))
	c.JSON(200, data)
}

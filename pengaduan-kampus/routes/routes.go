package routes

import (
	"github.com/gin-gonic/gin"
	"pengaduan-kampus/controllers"
	"pengaduan-kampus/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/submit", controllers.SubmitComplaint)
		auth.GET("/complaint/:id", controllers.StudentComplaintDetail)
	}

	admin := auth.Group("/admin")
	admin.Use(middlewares.AdminOnly())
	{
		admin.GET("/complaints", controllers.AllComplaints)
		admin.GET("/complaint/:id", controllers.ComplaintDetail)
		admin.POST("/complaint/:id/comment", controllers.AddComment)
		admin.GET("/complaint/:id/history", controllers.ActionHistory)
		admin.DELETE("/complaint/:id", controllers.DeleteComplaint)
		admin.PUT("/complaint/:id/status", controllers.UpdateComplaintStatus)
		admin.GET("/complaint/:id/status", controllers.ActionHistory)
	}
}

package main

import (
	"fmt"
	"log"

	"uas_cloudcomputing/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := setupRoutes()

	serverAddr := fmt.Sprintf("%s:%s",
		config.AppConfig.ServerHost,
		config.AppConfig.ServerPort,
	)

	log.Printf("Server starting on http://%s", serverAddr)

	if err := r.Run(serverAddr); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	// Static files
	r.Static("/static", "./web/static")
	r.Static("/uploads", "./uploads")

	// HTML routes (NO LOGIN)
	r.GET("/", func(c *gin.Context) {
		c.File("./template/complaints.html")
	})

	r.GET("/forgot-password", func(c *gin.Context) {
		c.File("./template/forgot_password.html")
	})

	r.GET("/admin/settings", func(c *gin.Context) {
		c.File("./template/settings.html")
	})

	r.GET("/reports", func(c *gin.Context) {
		c.File("./template/reports.html")
	})

	// API (dummy supaya compile)
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	return r
}

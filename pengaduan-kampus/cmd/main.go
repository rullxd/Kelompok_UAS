package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"pengaduan-kampus/config"
	"pengaduan-kampus/models"
	"pengaduan-kampus/routes"
)

func main() {
	// 1️⃣ Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 2️⃣ Load ENV ke struct
	config.LoadEnv()

	// 3️⃣ Connect Database
	config.ConnectDB()

	// 4️⃣ Auto migrate
	config.DB.AutoMigrate(
		&models.User{},
		&models.Complaint{},
		&models.Attachment{},
		&models.Comment{},
		&models.ActionLog{},
	)

	// 5️⃣ Setup server
	r := gin.Default()
	r.Static("/static", "./static")
	routes.SetupRoutes(r)

	// 6️⃣ Run server
	r.Run(":" + config.Env.AppPort)
}

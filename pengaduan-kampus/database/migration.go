package database

import (
	"pengaduan-kampus/config"
	"pengaduan-kampus/models"
)

func Migrate() {
	db := config.DB

	db.Migrator().AutoMigrate(
		&models.User{},
		&models.Complaint{},
	)

	db.Migrator().AutoMigrate(
		&models.Attachment{},
		&models.Comment{},
		&models.ActionLog{},
	)
}

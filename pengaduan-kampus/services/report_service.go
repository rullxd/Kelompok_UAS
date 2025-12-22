package services

import (
	"pengaduan-kampus/config"
	"pengaduan-kampus/models"
)

func GetAllComplaints() []models.Complaint {
	var data []models.Complaint
	config.DB.Find(&data)
	return data
}

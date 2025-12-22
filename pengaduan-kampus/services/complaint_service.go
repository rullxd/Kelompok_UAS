package services

import (
	"pengaduan-kampus/config"
	"pengaduan-kampus/models"
)

func CreateComplaint(userID uint, title, desc string) models.Complaint {
	complaint := models.Complaint{
		StudentID:  userID,
		Title:       title,
		Description: desc,
		Status:      "proses",
	}
	config.DB.Create(&complaint)
	return complaint
}

func GetComplaintDetail(id string) models.Complaint {
	var c models.Complaint
	config.DB.Preload("Attachments").Preload("Comments").First(&c, id)
	return c
}

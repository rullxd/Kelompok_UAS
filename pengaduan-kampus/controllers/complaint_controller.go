package controllers

import (
	"github.com/gin-gonic/gin"
	"pengaduan-kampus/config"
	"pengaduan-kampus/models"
	"pengaduan-kampus/services"
	"pengaduan-kampus/utils"
)

func SubmitComplaint(c *gin.Context) {
	userID := c.GetUint("user_id")
	title := c.PostForm("title")
	desc := c.PostForm("description")

	complaint := services.CreateComplaint(userID, title, desc)
	path := utils.SaveFile(c, "file")

	if path != "" {
		config.DB.Create(&models.Attachment{
			ComplaintID: complaint.ID,
			FilePath:    path,
		})
	}
	

	c.JSON(200, gin.H{"message": "Pengaduan terkirim"})
}
func DeleteComplaint(c *gin.Context) {
	id := c.Param("id")

	// hapus comment
	config.DB.Where("complaint_id = ?", id).Delete(&models.Comment{})

	// hapus attachment
	config.DB.Where("complaint_id = ?", id).Delete(&models.Attachment{})

	// hapus action log
	config.DB.Where("complaint_id = ?", id).Delete(&models.ActionLog{})

	// hapus complaint
	if err := config.DB.Delete(&models.Complaint{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal menghapus pengaduan"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Pengaduan berhasil dihapus",
	})
}

func UpdateComplaintStatus(c *gin.Context) {
	var input struct {
		Status string `json:"status"`
	}
	c.ShouldBindJSON(&input)

	id := c.Param("id")

	var complaint models.Complaint
	if err := config.DB.First(&complaint, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "Pengaduan tidak ditemukan"})
		return
	}

	oldStatus := complaint.Status

	// update status
	complaint.Status = input.Status
	config.DB.Save(&complaint)

	// simpan riwayat
	log := models.ActionLog{
		ComplaintID: complaint.ID,
		AdminID:     c.GetUint("user_id"),
		Action:      "Status diubah dari " + oldStatus + " ke " + input.Status,
	}

	config.DB.Create(&log)

	c.JSON(200, gin.H{
		"message": "Status pengaduan berhasil diperbarui",
	})
}


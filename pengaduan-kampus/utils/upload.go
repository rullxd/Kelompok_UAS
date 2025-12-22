package utils

import (
	"path/filepath"
	"github.com/gin-gonic/gin"
)

func SaveFile(c *gin.Context, field string) string {
	file, err := c.FormFile(field)
	if err != nil {
		return ""
	}
	path := "static/uploads/" + filepath.Base(file.Filename)
	c.SaveUploadedFile(file, path)
	return path
}

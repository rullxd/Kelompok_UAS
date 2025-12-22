package services

import (
	"errors"
	"pengaduan-kampus/config"
	"pengaduan-kampus/models"
)

func Login(email, password string) (string, error) {
	var user models.User
	config.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 || user.Password != password {
		return "", errors.New("login gagal")
	}
	return config.GenerateToken(user.ID, user.Role)
}

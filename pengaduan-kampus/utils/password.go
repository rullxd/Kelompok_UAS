package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	h, _ := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(h)
}

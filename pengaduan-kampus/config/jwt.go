package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY []byte

func InitJWT() {
	JWT_KEY = []byte(Env.JWTSecret)
}

func GenerateToken(id uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"role":    role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_KEY)
}

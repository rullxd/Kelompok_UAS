package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	DBDsn     string
	JWTSecret string
	AppPort   string
}

var Env AppConfig

func LoadEnv() {
	Env = AppConfig{
		AppPort:   getEnv("SERVER_PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "SECRET"),
		DBDsn: fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			getEnv("DB_USER", "root"),
			getEnv("DB_PASSWORD", ""),
			getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306"),
			getEnv("DB_NAME", ""),
		),
	}
}

func getEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

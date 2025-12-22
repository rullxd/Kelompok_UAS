package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open(Env.DBDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

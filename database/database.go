package database

import (
	"os"
	"sync"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	once sync.Once
)

func CreateConnection() {

	conf := dbConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if DB == nil {
		CreateConnection()
	}
	return DB
}
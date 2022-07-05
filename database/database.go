package database

import (
	"os"
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func getenv(key, fallback string) string {
	var (
		val     string
		isExist bool
	)
	val, isExist = os.LookupEnv(key)
	if !isExist {
		val = fallback
	}
	return val
}

func CreateConnection() {
	conf := dbConfig{
		User: getenv("DB_USER", "root"),
		Pass: getenv("DB_PASS", "1234567890"),
		Host: getenv("DB_HOST", "localhost"),
		Port: getenv("DB_PORT", "3306"),
		Name: getenv("DB_NAME", "employee_svc"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}

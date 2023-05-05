package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func init() {
	if dbConn != nil {
		return
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/vota?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost)
	dbConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if dbConn != nil {
		fmt.Printf("Database %s conected!\n", dbHost)
	}
}

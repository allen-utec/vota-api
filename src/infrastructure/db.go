package infrastructure

import (
	"fmt"

	"github.com/allen-utec/vota-api-polls/src/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbConn *gorm.DB

func init() {
	if dbConn != nil {
		return
	}

	dbUser := getEnv("DB_USER", "")
	dbPass := getEnv("DB_PASS", "")
	dbHost := getEnv("DB_HOST", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/vota?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost)
	dbConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	dbConn.AutoMigrate(&domain.Alternative{}, &domain.Poll{}, &domain.Vote{})

	if dbConn != nil {
		fmt.Printf("Database %s conected!\n", dbHost)
	}
}

package infrastructure

import (
	"fmt"
	"os"

	"github.com/allen-utec/vota-api/src/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMysqlConnection() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/vota?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost)
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = dbConn.AutoMigrate(&domain.User{}, &domain.Poll{}, &domain.PollOption{})
	return dbConn, err
}

func InitRepositories(dbConn *gorm.DB) {
	PollRepositoryInstance = &PollRepository{dbConn}
	UserRepositoryInstance = &UserRepository{dbConn}
}

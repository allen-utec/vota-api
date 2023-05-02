package infrastructure

import (
	"github.com/allen-utec/vota-api/src/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbConn *gorm.DB
}

var UserRepositoryInstance *UserRepository

func InitUserRepository(dbConn *gorm.DB) {
	if UserRepositoryInstance == nil {
		UserRepositoryInstance = &UserRepository{dbConn}
	}
}

func (r *UserRepository) Create(user domain.User) (domain.User, error) {
	result := r.dbConn.FirstOrCreate(&user)

	return user, result.Error
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	result := r.dbConn.Find(&users)

	return users, result.Error
}

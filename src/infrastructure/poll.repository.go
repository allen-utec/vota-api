package infrastructure

import (
	"github.com/allen-utec/vota-api/src/domain"
	"gorm.io/gorm"
)

type PollRepository struct {
	dbConn *gorm.DB
}

var PollRepositoryInstance *PollRepository

func InitPollRepository(dbConn *gorm.DB) {
	if PollRepositoryInstance == nil {
		PollRepositoryInstance = &PollRepository{dbConn}
	}
}

func (r *PollRepository) Create(poll domain.Poll) (domain.Poll, error) {
	result := r.dbConn.Create(&poll)

	return poll, result.Error
}

func (r *PollRepository) GetAll() ([]domain.Poll, error) {
	var polls []domain.Poll
	result := r.dbConn.Find(&polls)

	return polls, result.Error
}

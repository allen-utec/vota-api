package infrastructure

import (
	"github.com/allen-utec/vota-api/src/domain"
)

type PollRepository struct{}

func (r *PollRepository) Create(poll domain.Poll) (domain.Poll, error) {
	result := dbConn.Create(&poll)

	return poll, result.Error
}

func (r *PollRepository) GetAll() ([]domain.Poll, error) {
	var polls []domain.Poll
	result := dbConn.Find(&polls)

	return polls, result.Error
}

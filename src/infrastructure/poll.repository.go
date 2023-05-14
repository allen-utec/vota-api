package infrastructure

import (
	"github.com/allen-utec/vota-api-polls/src/domain"
)

type PollRepository struct{}

func (r *PollRepository) Create(poll domain.Poll) (domain.Poll, error) {
	result := dbConn.Create(&poll)

	return poll, result.Error
}

func (r *PollRepository) GetAll() ([]domain.Poll, error) {
	var polls []domain.Poll
	result := dbConn.Preload("Alternatives").Find(&polls)

	return polls, result.Error
}

func (r *PollRepository) GetByCode(code string) (domain.Poll, error) {
	var poll domain.Poll
	result := dbConn.Where(&domain.Poll{Code: code}).Preload("Alternatives").First(&poll)

	return poll, result.Error
}

func (r *PollRepository) GetResults(id uint) (domain.Poll, error) {
	var poll domain.Poll
	result := dbConn.Preload("Alternatives").Preload("Votes").First(&poll, id)
	return poll, result.Error
}

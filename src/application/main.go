package application

import "github.com/allen-utec/vota-api/src/domain"

/* Use Cases */

type PollUseCases struct {
	repository domain.PollRepository
}

func (uc *PollUseCases) GetAllPoll() ([]domain.Poll, error) {
	return uc.repository.GetAll()
}

func (uc *PollUseCases) CreatePoll(poll domain.Poll) (domain.Poll, error) {
	return uc.repository.Create(poll)
}

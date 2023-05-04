package application

import "github.com/allen-utec/vota-api/src/domain"

type pollService struct {
	repository domain.PollRepository
}

/* Use Cases */

func (s *pollService) GetAllUseCase() ([]domain.Poll, error) {
	return s.repository.GetAll()
}

func (s *pollService) CreateUseCase(poll domain.Poll) (domain.Poll, error) {
	return s.repository.Create(poll)
}

var PollService pollService

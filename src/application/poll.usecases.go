package application

import "github.com/allen-utec/vota-api/src/domain"

type pollService struct {
	repository domain.PollRepository
}

var PollService pollService

func (s *pollService) Init(r domain.PollRepository) {
	s.repository = r
}

/* Use Cases */

func (s *pollService) GetAllUseCase() ([]domain.Poll, error) {
	return s.repository.GetAll()
}

func (s *pollService) GetByCodeUseCase(code string) (domain.Poll, error) {
	return s.repository.GetByCode(code)
}

func (s *pollService) CreateUseCase(poll domain.Poll) (domain.Poll, error) {
	return s.repository.Create(poll)
}

package application

import "github.com/allen-utec/vota-api/src/domain"

type userService struct {
	repository domain.UserRepository
}

var UserService userService

func (s *userService) Init(r domain.UserRepository) {
	s.repository = r
}

/* Use Cases */

func (s *userService) GetAllUseCase() ([]domain.User, error) {
	return s.repository.GetAll()
}

func (s *userService) CreateUseCase(user domain.User) (domain.User, error) {
	return s.repository.Create(user)
}

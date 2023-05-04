package application

import "github.com/allen-utec/vota-api/src/domain"

func InitServices(pollRepository domain.PollRepository, userRepository domain.UserRepository) {
	PollService.repository = pollRepository
	UserService.repository = userRepository
}

package application

import "github.com/allen-utec/vota-api/src/domain"

type voteService struct {
	repository domain.VoteRepository
}

var VoteService voteService

func (s *voteService) Init(r domain.VoteRepository) {
	s.repository = r
}

/* Use Cases */

func (s *voteService) CreateUseCase(vote domain.Vote) (domain.Vote, error) {
	return s.repository.Create(vote)
}

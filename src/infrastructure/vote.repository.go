package infrastructure

import (
	"github.com/allen-utec/vota-api/src/domain"
)

type VoteRepository struct{}

func (r *VoteRepository) Create(vote domain.Vote) (domain.Vote, error) {
	result := dbConn.Create(&vote)

	return vote, result.Error
}

package domain

import (
	"gorm.io/gorm"
)

/* Entities */

type Poll struct {
	gorm.Model
	Question     string
	UserID       uint
	Alternatives []Alternative
	Code         string
	Votes        []Vote
}

type Alternative struct {
	gorm.Model
	Text   string
	PollID uint
}

type Vote struct {
	gorm.Model
	PollID        uint
	AlternativeID uint
	UserID        uint
}

type PollRepository interface {
	Create(poll Poll) (Poll, error)
	GetAll() ([]Poll, error)
	GetByCode(code string) (Poll, error)
	GetResults(id uint) (Poll, error)
}

type VoteRepository interface {
	Create(vote Vote) (Vote, error)
}

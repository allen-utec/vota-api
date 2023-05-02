package domain

import (
	"gorm.io/gorm"
)

/* Entities */

type User struct {
	gorm.Model
	Username string
	Polls    []Poll
}

type Poll struct {
	gorm.Model
	Question    string
	UserID      uint
	PollOptions []PollOption
}

type PollOption struct {
	gorm.Model
	Text   string
	PollID uint
}

type PollRepository interface {
	Create(poll Poll) (Poll, error)
	GetAll() ([]Poll, error)
}

package domain

import (
	"gorm.io/gorm"
)

/* Entities */

type User struct {
	gorm.Model
	Nickname string
	Polls    []Poll
}

type Poll struct {
	gorm.Model
	Question     string
	UserID       uint
	Alternatives []Alternative
	Code         string
}

type Alternative struct {
	gorm.Model
	Text   string
	PollID uint
}

type PollRepository interface {
	Create(poll Poll) (Poll, error)
	GetAll() ([]Poll, error)
	GetByCode(code string) (Poll, error)
}

type UserRepository interface {
	Create(user User) (User, error)
	GetAll() ([]User, error)
}

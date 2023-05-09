package infrastructure

import "github.com/allen-utec/vota-api/src/domain"

type PollVM struct {
	ID           uint            `json:"id"`
	Question     string          `json:"question"`
	Alternatives []AlternativeVM `json:"alternatives"`
	UserID       uint            `json:"user_id"`
	Code         string          `json:"code"`
}

type AlternativeVM struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

type UserVM struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
}

func formatPoll(poll domain.Poll) PollVM {
	alternatives := make([]AlternativeVM, len(poll.Alternatives))

	for i, e := range poll.Alternatives {
		alternatives[i] = AlternativeVM{ID: e.ID, Text: e.Text}
	}

	pollVM := PollVM{
		ID:           poll.ID,
		Question:     poll.Question,
		Alternatives: alternatives,
		UserID:       poll.UserID,
		Code:         poll.Code,
	}
	return pollVM
}

func formatUser(user domain.User) UserVM {
	userVM := UserVM{
		ID:       user.ID,
		Nickname: user.Nickname,
	}
	return userVM
}

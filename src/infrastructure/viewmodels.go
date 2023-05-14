package infrastructure

import "github.com/allen-utec/vota-api-polls/src/domain"

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

type VoteVM struct {
	AlternativeID uint `json:"alternative_id"`
	UserID        uint `json:"user_id"`
}

type PollResultsVM struct {
	Question     string          `json:"question"`
	Alternatives []AlternativeVM `json:"alternatives"`
	Votes        []uint          `json:"votes"`
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

func formatPollResults(poll domain.Poll) PollResultsVM {
	pollVM := formatPoll(poll)

	votes := make([]uint, len(poll.Votes))
	for i, e := range poll.Votes {
		votes[i] = e.AlternativeID
	}

	pollResultsVM := PollResultsVM{
		Question:     pollVM.Question,
		Alternatives: pollVM.Alternatives,
		Votes:        votes,
	}
	return pollResultsVM
}

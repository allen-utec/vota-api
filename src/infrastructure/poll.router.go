package infrastructure

import (
	"net/http"

	"github.com/allen-utec/vota-api/src/application"
	"github.com/allen-utec/vota-api/src/domain"
	"github.com/gin-gonic/gin"
)

type poll struct {
	ID       string       `json:"id"`
	Question string       `json:"question"`
	Options  []pollOption `json:"options"`
	UserID   uint         `json:"user_id"`
}

type pollOption struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func CreatePoll(ctx *gin.Context) {
	var payload poll

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	newPoll := domain.Poll{
		Question:    payload.Question,
		PollOptions: make([]domain.PollOption, len(payload.Options)),
		UserID:      payload.UserID,
	}

	for i, e := range payload.Options {
		newPoll.PollOptions[i] = domain.PollOption{Text: e.Text}
	}

	poll, err := application.PollService.CreateUseCase(newPoll)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, poll)
}

func GetAllPolls(ctx *gin.Context) {
	polls, err := application.PollService.GetAllUseCase()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, polls)
}

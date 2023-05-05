package infrastructure

import (
	"net/http"

	"github.com/allen-utec/vota-api/src/application"
	"github.com/allen-utec/vota-api/src/domain"
	"github.com/gin-gonic/gin"
)

func CreatePollHandler(ctx *gin.Context) {
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

func GetAllPollsHandler(ctx *gin.Context) {
	polls, err := application.PollService.GetAllUseCase()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, polls)
}

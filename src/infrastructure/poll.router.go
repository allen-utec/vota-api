package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/allen-utec/vota-api-polls/src/application"
	"github.com/allen-utec/vota-api-polls/src/domain"
	"github.com/gin-gonic/gin"
)

func CreatePollHandler(ctx *gin.Context) {
	var payload PollVM

	if err := ctx.BindJSON(&payload); err != nil {
		reponseError(ctx, err)
		return
	}

	newPoll := domain.Poll{
		Question:     payload.Question,
		Alternatives: make([]domain.Alternative, len(payload.Alternatives)),
		UserID:       payload.UserID,
		Code:         generateID(),
	}

	for i, e := range payload.Alternatives {
		newPoll.Alternatives[i] = domain.Alternative{Text: e.Text}
	}

	poll, err := application.PollService.CreateUseCase(newPoll)
	if err != nil {
		reponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, formatPoll(poll))
}

func GetAllPollsHandler(ctx *gin.Context) {
	polls, err := application.PollService.GetAllUseCase()
	if err != nil {
		reponseError(ctx, err)
		return
	}

	pollsVM := make([]PollVM, len(polls))
	for i, e := range polls {
		pollsVM[i] = formatPoll(e)
	}

	ctx.JSON(http.StatusOK, pollsVM)
}

func GetPollByCodeHandler(ctx *gin.Context) {
	poll, err := application.PollService.GetByCodeUseCase(ctx.Param("code"))
	if err != nil {
		reponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, formatPoll(poll))
}

func CreateVoteHandler(ctx *gin.Context) {
	var payload VoteVM

	if err := ctx.BindJSON(&payload); err != nil {
		reponseError(ctx, err)
		return
	}

	pollId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		reponseError(ctx, err)
		return
	}

	newVote := domain.Vote{
		PollID:        uint(pollId),
		AlternativeID: payload.AlternativeID,
		UserID:        payload.UserID,
	}

	_, err = application.VoteService.CreateUseCase(newVote)

	if err != nil {
		reponseError(ctx, err)
		return
	}

	pollResults, err := application.PollService.GetResultsUseCase(uint(pollId))

	ctx.JSON(http.StatusCreated, formatPollResults(pollResults))
}

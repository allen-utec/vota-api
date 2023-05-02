package infrastructure

import (
	"net/http"

	"github.com/allen-utec/vota-api/src/domain"
	"github.com/gin-gonic/gin"
)

type poll struct {
	ID       string       `json:"id"`
	Question string       `json:"question"`
	Options  []pollOption `json:"options"`
	User     user         `json:"user"`
}

type pollOption struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func CreatePoll(ctx *gin.Context) {
	var newPoll domain.Poll

	if err := ctx.BindJSON(&newPoll); err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	poll, err := PollRepositoryInstance.Create(newPoll)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, poll)
}

func GetAllPolls(ctx *gin.Context) {
	polls, err := PollRepositoryInstance.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, polls)
}

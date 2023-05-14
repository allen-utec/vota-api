package infrastructure

import (
	"fmt"

	"github.com/allen-utec/vota-api-polls/src/application"
	"github.com/gin-gonic/gin"
)

func init() {
	application.PollService.Init(&PollRepository{})
	application.VoteService.Init(&VoteRepository{})
}

func InitApi() {
	router := gin.Default()
	router.Use(corsMiddleware())

	router.GET("/polls", GetAllPollsHandler)
	router.POST("/polls", CreatePollHandler)
	router.GET("/polls/:code", GetPollByCodeHandler)
	router.POST("/polls/:id", CreateVoteHandler)

	router.Run(fmt.Sprintf(":%s", getEnv("PORT", "8080")))
}

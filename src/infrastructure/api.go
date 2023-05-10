package infrastructure

import (
	"fmt"

	"github.com/allen-utec/vota-api/src/application"
	"github.com/gin-gonic/gin"
)

func init() {
	application.PollService.Init(&PollRepository{})
	application.UserService.Init(&UserRepository{})
}

func InitApi() {
	router := gin.Default()
	router.Use(corsMiddleware())

	router.GET("/polls", GetAllPollsHandler)
	router.GET("/polls/:code", GetPollByCodeHandler)
	router.POST("/polls", CreatePollHandler)

	router.GET("/users", GetAllUsersHandler)
	router.POST("/users", CreateUserHandler)

	router.Run(fmt.Sprintf(":%s", getEnv("PORT", "8080")))
}

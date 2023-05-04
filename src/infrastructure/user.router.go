package infrastructure

import (
	"net/http"

	"github.com/allen-utec/vota-api/src/application"
	"github.com/allen-utec/vota-api/src/domain"
	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func CreateUser(ctx *gin.Context) {
	var payload user

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	newUser := domain.User{
		Username: payload.Username,
	}

	user, err := application.UserService.CreateUseCase(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := application.UserService.GetAllUseCase()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

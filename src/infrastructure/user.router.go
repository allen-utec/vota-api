package infrastructure

import (
	"net/http"

	"github.com/allen-utec/vota-api/src/application"
	"github.com/allen-utec/vota-api/src/domain"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	var payload UserVM

	if err := ctx.BindJSON(&payload); err != nil {
		reponseError(ctx, err)
		return
	}

	newUser := domain.User{
		Nickname: payload.Nickname,
	}

	user, err := application.UserService.CreateUseCase(newUser)
	if err != nil {
		reponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, formatUser(user))
}

func GetAllUsersHandler(ctx *gin.Context) {
	users, err := application.UserService.GetAllUseCase()
	if err != nil {
		reponseError(ctx, err)
		return
	}

	usersMV := make([]UserVM, len(users))
	for i, e := range users {
		usersMV[i] = formatUser(e)
	}
	ctx.JSON(http.StatusOK, usersMV)
}

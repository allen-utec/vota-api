package infrastructure

import (
	"net/http"

	"github.com/allen-utec/vota-api/src/domain"
	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func CreateUser(ctx *gin.Context) {
	var newUser domain.User

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	user, err := UserRepositoryInstance.Create(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := UserRepositoryInstance.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &responseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

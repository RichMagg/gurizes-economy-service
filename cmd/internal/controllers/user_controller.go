package controllers

import (
	"net/http"

	"github.com/RichMagg/gurizes-economy-service/cmd/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	users, err := u.userService.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	user, err := u.userService.CreateUser()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, user)
}

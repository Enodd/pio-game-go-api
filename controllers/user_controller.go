package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
}

type userController struct{}

func NewUserController() *userController {
	return &userController{}
}

// GetUser returns user with given id
//
//	@Summary		returns user
//	@Description	returns user with given id
//	@Tags			users
//	@Produce		plain
//	@Param			id	path		int	true	"id of user"
//	@Success		200	{string}	string
//	@Router			/users/{id} [get]
func (uc *userController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(errors.New("id wasnt provided"))
	}
	if aid == 0 {
		ctx.Error(errors.New("cannot get users"))
	}
	response := "Jeff"
	ctx.JSON(http.StatusOK, response)
}

// DeleteUser deletes user
//
//	@Summary		Deletes user
//	@Description	Deletes user of a given id
//	@Tags			users
//	@Produce		plain
//	@Param			id	path		int	true	"id of user"
//	@Success		200	{string} string
//	@Router			/users/{id} [delete]
func (uc *userController) DeleteUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Endpoint to be implemented")
}

// UpdateUser updates user info
//
//	@Summary		updates user
//	@Description	updates user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user formData User true "user data"
//	@Success		200	{string}	string
//	@Router			/users/{id} [put]
func (uc *userController) UpdateUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Endpoint to be implemented")
}

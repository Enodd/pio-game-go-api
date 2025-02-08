package controllers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"pioApi/config"
	_ "pioApi/ent"
	"pioApi/models"
	services "pioApi/services"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service *services.UserService
}

func NewUserController() *userController {
	return &userController{
		service: nil,
	}
}

func (uc *userController) SetService(service *services.UserService) {
	uc.service = service
}

func (uc *userController) InitRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", uc.GetUsers)
		users.GET("/:id", uc.GetUser)
		users.POST("/register", uc.CreateUser)
		users.PATCH("/:id", uc.UpdateUser)
		users.DELETE("/:id", uc.DeleteUser)
	}
}

// GetUsers returns list of users
//
//	@Summary		returns list of users
//	@Description	returns list of all users in database
//	@Tags			users
//	@Produce		json
//	@Success		200	{array}	ent.User
//	@Router			/users/ [get]
func (uc *userController) GetUsers(ctx *gin.Context) {
	context := context.Background()

	users, err := uc.service.GetUsers(context, config.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
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

// CreateUser returns user with given id
//
//	@Summary		returns user
//	@Description	returns user with given id
//	@Tags			users
//	@Produce		plain
//	@Param			id	path		int	true	"id of user"
//	@Success		200	{string}	string
//	@Router			/users/{id} [post]
func (uc *userController) CreateUser(ctx *gin.Context) {
	context := context.Background()
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := uc.service.CreateUser(user, context, config.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "User created successfully")
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
//	@Param			user formData string true "user data"
//	@Success		200	{string}	string
//	@Router			/users/{id} [patch]
func (uc *userController) UpdateUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Endpoint to be implemented")
}

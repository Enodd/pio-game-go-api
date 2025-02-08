package services

import (
	"context"
	"pioApi/ent"
	"pioApi/models"
	repositories "pioApi/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) SetRepository(repository *repositories.UserRepository) {
	us.repository = repository
}

func (us *UserService) CreateUser(user models.User, ctx context.Context, client *ent.Client) (bool, error) {
	success, err := us.repository.CreateUser(user, ctx, client)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (us *UserService) GetUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	users, err := us.repository.GetUsers(ctx, client)
	if err != nil {
		return nil, err
	}
	return users, nil
}

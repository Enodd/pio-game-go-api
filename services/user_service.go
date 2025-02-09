package services

import (
	"context"
	"pioApi/ent"
	"pioApi/models"
	repositories "pioApi/repositories"

	"github.com/google/uuid"
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
	err := us.repository.CreateUser(user, ctx, client)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (us *UserService) GetUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	users, err := us.repository.GetUsers(ctx, client)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetUser(id uuid.UUID, ctx context.Context, client *ent.Client) (*ent.User, error) {
	user, err := us.repository.GetUser(id, ctx, client)

	if err != nil {
		return nil, err
	}

	return user, nil
}

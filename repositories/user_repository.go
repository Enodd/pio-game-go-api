package repositories

import (
	"context"
	"pioApi/ent"
	"pioApi/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(user models.User, ctx context.Context, client *ent.Client) (bool, error) {
	err := client.User.
		Create().
		SetUsername(user.Name).
		SetPassword(user.Password).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ur *UserRepository) GetUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	users, err := client.User.Query().All(ctx)

	if err != nil {
		return nil, err
	}
	return users, nil
}

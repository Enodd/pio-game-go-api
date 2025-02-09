package repositories

import (
	"context"
	"pioApi/ent"
	"pioApi/ent/user"
	"pioApi/models"

	"github.com/google/uuid"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(user models.User, ctx context.Context, client *ent.Client) error {
	return client.User.
		Create().
		SetUsername(user.Name).
		SetPassword(user.Password).
		Exec(ctx)
}

func (ur *UserRepository) GetUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	return client.User.Query().All(ctx)
}

func (ur *UserRepository) GetUser(id uuid.UUID, ctx context.Context, client *ent.Client) (*ent.User, error) {
	return client.User.Query().Where(user.ID(id)).First(ctx)
}

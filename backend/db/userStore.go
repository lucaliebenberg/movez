package db

import (
	"context"

	"github.com/lucaliebenberg/movez/backend/types"
)

type SqlUserStore struct {}

type Map map[string]any

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUserByEmail(context.Context, string) (*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error
}

func (s *SqlUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	return &user, nil
}
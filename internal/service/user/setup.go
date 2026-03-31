package user

import (
	"backend/internal/model"
	"context"
)

type userStorage interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
}

type Service struct {
	user userStorage
}

func New(userStorage userStorage) *Service {
	return &Service{
		user: userStorage,
	}
}

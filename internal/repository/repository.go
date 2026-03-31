package repository

import (
	"backend/internal/model"
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}

type TestRepository interface {
	//getAllTest(ctx context.Context) []model.TestFull
	GetTestById(ctx context.Context, id int) (model.TestFull, error)
	GetAvailableTests(ctx context.Context, user_id int) ([]model.TestFull,error)
}

type LoginRepository interface {
	LogIn(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error)
}

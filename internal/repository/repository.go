package repository

import(
	"backend/internal/model"
	"context"
)

type UserRepository interface {
    GetByID(ctx context.Context, id int) (*model.User, error)
    GetAll(ctx context.Context) ([]*model.User, error)
}

package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"fmt"
)

type UserService interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetByID(ctx context.Context, id int) (*model.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	return s.userRepo.GetByID(ctx, id)
}

func (s *userService) GetAll(ctx context.Context) ([]*model.User, error) {
	return s.userRepo.GetAll(ctx)
}

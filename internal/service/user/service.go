package user

import (
	"backend/internal/model"
	"context"
)

func (s *Service) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.user.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, model.ErrUserNotFound
	}

	return user, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*model.User, error) {
	return s.user.GetAllUsers(ctx)
}

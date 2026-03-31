package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"errors"
)

type LoginService interface {
	LogIn(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error)
}
type loginService struct {
	loginRepo repository.LoginRepository
}

func NewLoginService(loginRepo repository.LoginRepository) LoginService {
	return &loginService{
		loginRepo: loginRepo,
	}
}

func (r *loginService) LogIn(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {

	resp, err := r.loginRepo.LogIn(ctx, req)

	if err != nil {
		return resp, err
	}

	if isEmpty(resp) {
		return resp, errors.New("User not found")
	}

	return resp, nil
}

func isEmpty(res model.LoginResponse) bool {
	return (res.Id == 0 &&
		res.Name == "" &&
		res.Role == "")
}

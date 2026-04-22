package loginservice

import (
	"backend/internal/model"
	"context"
	"errors"
)

func (s *LoginService) SignIn(ctx context.Context, loginRequest model.LoginRequest) (model.LoginResponse, error) {

	response, err := s.repository.SignIn(ctx, loginRequest)
	//TODO обработать ошибки (все в проектк)
	//if err != nil {
	//	return response, err
	//}

	//TODO вынести в константы ошибку
	if response.Id == 0 && response.Name == "" && response.Role == "" {
		return response, errors.New("User not found")
	}

	return response, err
}

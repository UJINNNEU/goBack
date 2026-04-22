package loginservice

import (
	"backend/internal/repository/storage/login"
)
type LoginService struct{
	repository *login.LoginRepository
}

func NewService(loginRepository *login.LoginRepository ) *LoginService{
	return &LoginService{
		repository: loginRepository,
	}
}
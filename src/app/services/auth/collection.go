package auth_service

import (
	"test-kreditplus/src/app/entities"
)

type AuthService struct {
	authRepo AuthRepository
}

func NewAuthService(authRepo AuthRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

type AuthRepository interface {
	Create(request entities.Auth) (response *entities.Auth, err error)
	Login(username string, password string) (token string, err error)
	Get(id uint) (response entities.Auth, err error)
}

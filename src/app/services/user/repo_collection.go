package user_service

import (
	"kredit_plus/src/app/entities"
)

type UserService struct {
	userRepo UserRepository
}

func NewMerchantService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type UserRepository interface {
	Create(request entities.User) (response *entities.User, err error)
}

package handlers

import (
	"test-kreditplus/src/contract"
	"test-kreditplus/src/entities"
)

type AuthService interface {
	Create(request contract.RegisterInput) (response *contract.RegisterOutput, err error)
	Login(request contract.LoginInput) (token string, err error)
	Get(id uint) (response entities.Auth, err error)
}

type ProfileService interface {
	Create(request contract.ProfileInput) (response *contract.ProfileInput, err error)
}

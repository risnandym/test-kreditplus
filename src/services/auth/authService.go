package auth_service

import (
	"test-kreditplus/src/contract"
	"test-kreditplus/src/entities"
	"time"
)

func (a AuthService) Create(request contract.RegisterInput) (response *contract.RegisterOutput, err error) {

	auth := entities.Auth{}
	auth.Email = request.Email
	auth.Phone = request.Phone
	auth.Password = request.Password
	auth.LastLogin = time.Now()

	result, err := a.authRepo.Create(auth)
	if err != nil {
		return nil, err
	}

	response = &contract.RegisterOutput{}
	response.Email = result.Email
	response.LastLogin = result.LastLogin
	response.Phone = result.Phone
	response.UUID = result.UUID

	return
}

func (a AuthService) Login(request contract.LoginInput) (token string, err error) {

	auth := entities.Auth{}
	auth.Email = request.Email
	auth.Password = request.Password

	token, err = a.authRepo.Login(auth.Email, auth.Password)

	return
}

func (a AuthService) Get(id uint) (response entities.Auth, err error) {
	return a.authRepo.Get(id)
}

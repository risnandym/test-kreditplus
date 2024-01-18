package user_service

import (
	"kredit_plus/src/app/contract"
	"kredit_plus/src/app/entities"
	"time"
)

func (u UserService) Create(input contract.RegisterInput) {

	user := entities.User{}
	user.Email = input.Email
	user.Email = input.Email
	user.Password = input.Password
	user.LastLogin = time.Now()

	u.userRepo.Create(user)

}

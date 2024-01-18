package handlers

import (
	"kredit_plus/src/app/contract"
)

type UserService interface {
	Create(input contract.RegisterInput)
}

package profile_service

import (
	"test-kreditplus/src/app/entities"
)

type ProfileService struct {
	profileRepo ProfileRepository
}

func NewProfileService(profileRepo ProfileRepository) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
	}
}

type ProfileRepository interface {
	Create(request entities.Profile) (response *entities.Profile, err error)
}

// type ProfileRepository interface {
// 	Create(request entities.Auth) (response *entities.Auth, err error)
// 	Login(username string, password string) (token string, err error)
// }

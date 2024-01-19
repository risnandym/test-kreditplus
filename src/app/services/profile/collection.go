package profile_service

import (
	"test-kreditplus/src/app/entities"
)

type ProfileService struct {
	profileRepo ProfileRepository
	limitRepo   LimitRepository
}

func NewProfileService(profileRepo ProfileRepository, limitRepo LimitRepository) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
		limitRepo:   limitRepo,
	}
}

type ProfileRepository interface {
	Create(request entities.Profile) (response *entities.Profile, err error)
}

type LimitRepository interface {
	Create(request entities.Limit) (response *entities.Limit, err error)
}

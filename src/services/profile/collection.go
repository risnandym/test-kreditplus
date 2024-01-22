package profile_service

import (
	"test-kreditplus/src/entities"

	"gorm.io/gorm"
)

type ProfileService struct {
	db          *gorm.DB
	profileRepo ProfileRepository
	limitRepo   LimitRepository
}

func NewProfileService(db *gorm.DB, profileRepo ProfileRepository, limitRepo LimitRepository) *ProfileService {
	return &ProfileService{
		db:          db,
		profileRepo: profileRepo,
		limitRepo:   limitRepo,
	}
}

type ProfileRepository interface {
	Create(db *gorm.DB, request *entities.Profile) (response *entities.Profile, err error)
}

type LimitRepository interface {
	Create(db *gorm.DB, request *entities.Limit) (response *entities.Limit, err error)
}

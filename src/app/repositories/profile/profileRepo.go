package profile_repo

import (
	"test-kreditplus/src/app/entities"
	"time"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) (*ProfileRepository, error) {
	return &ProfileRepository{
		db: db,
	}, nil
}

func (p ProfileRepository) Create(request entities.Profile) (response *entities.Profile, err error) {

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = p.db.Create(&request).Error; err != nil {
		return
	}

	return
}

package limit_repo

import (
	"test-kreditplus/src/entities"
	"time"

	"gorm.io/gorm"
)

type LimitRepository struct {
	db *gorm.DB
}

func NewLimitRepository(db *gorm.DB) (*LimitRepository, error) {
	return &LimitRepository{
		db: db,
	}, nil
}

func (l LimitRepository) Create(request entities.Limit) (response *entities.Limit, err error) {

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = l.db.Create(&request).Error; err != nil {
		return
	}

	return
}

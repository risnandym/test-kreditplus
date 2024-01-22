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

func (l LimitRepository) Create(db *gorm.DB, request *entities.Limit) (response *entities.Limit, err error) {

	if db != nil {
		l.db = db
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = l.db.Create(&request).Error; err != nil {
		return
	}

	response = request
	return
}

func (l LimitRepository) Update(db *gorm.DB, request *entities.Limit) (response *entities.Limit, err error) {

	if db != nil {
		l.db = db
	}

	request.UpdatedAt = time.Now()
	if err = l.db.Save(&request).Error; err != nil {
		return
	}

	response = request
	return
}

func (l LimitRepository) Get(id uint) (response *entities.Limit, err error) {

	if err = l.db.Where("auth_id = ?", id).First(&response).Error; err != nil {
		return
	}

	return
}

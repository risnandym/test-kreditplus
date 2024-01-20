package credit_repo

import (
	"test-kreditplus/src/entities"
	"time"

	"gorm.io/gorm"
)

type CreditRepository struct {
	db *gorm.DB
}

func NewCreditRepository(db *gorm.DB) (*CreditRepository, error) {
	return &CreditRepository{
		db: db,
	}, nil
}

func (t CreditRepository) Create(db *gorm.DB, request *entities.CreditTransaction) (response *entities.CreditTransaction, err error) {

	if db != nil {
		t.db = db
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = t.db.Create(&request).Error; err != nil {
		return
	}

	return
}

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

func (c CreditRepository) Create(db *gorm.DB, request *entities.Credit) (response *entities.Credit, err error) {

	if db != nil {
		c.db = db
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = c.db.Create(&request).Error; err != nil {
		return
	}

	response = request
	return
}

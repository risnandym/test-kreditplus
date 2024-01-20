package transaction_service

import (
	"test-kreditplus/src/entities"

	"gorm.io/gorm"
)

type TransactionService struct {
	db         *gorm.DB
	creditRepo CreditRepository
	limitRepo  LimitRepository
}

func NewTransactionService(db *gorm.DB, creditRepo CreditRepository, limitRepo LimitRepository) *TransactionService {
	return &TransactionService{
		db:         db,
		creditRepo: creditRepo,
		limitRepo:  limitRepo,
	}
}

type CreditRepository interface {
	Create(db *gorm.DB, request *entities.CreditTransaction) (response *entities.CreditTransaction, err error)
}

type LimitRepository interface {
	Get(id uint) (response *entities.Limit, err error)
	Update(db *gorm.DB, request *entities.Limit) (response *entities.Limit, err error)
}

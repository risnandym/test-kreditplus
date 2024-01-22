package transaction_service

import (
	"test-kreditplus/src/entities"

	"gorm.io/gorm"
)

type TransactionService struct {
	db         *gorm.DB
	creditRepo CreditRepository
	debitRepo  DebitRepository
	limitRepo  LimitRepository
	assetRepo  AssetRepository
}

func NewTransactionService(db *gorm.DB, creditRepo CreditRepository, debitRepo DebitRepository, limitRepo LimitRepository, assetRepo AssetRepository) *TransactionService {
	return &TransactionService{
		db:         db,
		creditRepo: creditRepo,
		debitRepo:  debitRepo,
		limitRepo:  limitRepo,
		assetRepo:  assetRepo,
	}
}

type LimitRepository interface {
	Get(id uint) (response *entities.Limit, err error)
	Update(db *gorm.DB, request *entities.Limit) (response *entities.Limit, err error)
}

type CreditRepository interface {
	Create(db *gorm.DB, request *entities.Credit) (response *entities.Credit, err error)
}

type DebitRepository interface {
	Get(month string, year int) (response *entities.Debit, err error)
	CreateInBatches(db *gorm.DB, request []entities.Debit, month string) (response []entities.Debit, err error)
	Create(db *gorm.DB, request *entities.Debit) (response *entities.Debit, err error)
	Update(db *gorm.DB, request *entities.Debit) (response *entities.Debit, err error)
}

type AssetRepository interface {
	Create(db *gorm.DB, request *entities.Asset) (response *entities.Asset, err error)
	CreateInBatches(db *gorm.DB, request []*entities.Asset) (response []*entities.Asset, err error)
}

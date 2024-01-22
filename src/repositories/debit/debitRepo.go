package debit_repo

import (
	"test-kreditplus/src/entities"
	"time"

	"gorm.io/gorm"
)

type DebitRepository struct {
	db *gorm.DB
}

func NewDebitRepository(db *gorm.DB) (*DebitRepository, error) {
	return &DebitRepository{
		db: db,
	}, nil
}

func (d DebitRepository) Get(month string, year int) (response *entities.Debit, err error) {

	if err = d.db.Where("month = ? AND year = ?", month, year).First(&response).Error; err != nil {
		return
	}

	return
}

func (d DebitRepository) Create(db *gorm.DB, request *entities.Debit) (response *entities.Debit, err error) {

	if db != nil {
		d.db = db
	}

	// request.UUID = uuid.New()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = d.db.Create(request).Error; err != nil {
		return
	}

	response = request

	return
}

func (d DebitRepository) CreateInBatches(db *gorm.DB, request []entities.Debit, month string) (response []entities.Debit, err error) {

	if db != nil {
		d.db = db
	}

	for idx, val := range request {
		parsedMonth, _ := time.Parse("January", month)
		now := parsedMonth.AddDate(0, idx, 0)
		monthstring := now.Month().String()
		val.Month = monthstring
		val.CreatedAt = time.Now()
		val.UpdatedAt = time.Now()
	}

	if err = d.db.CreateInBatches(request, len(request)).Error; err != nil {
		return
	}

	response = request

	return
}

func (d DebitRepository) Update(db *gorm.DB, request *entities.Debit) (response *entities.Debit, err error) {

	if db != nil {
		d.db = db
	}

	request.UpdatedAt = time.Now()
	if err = d.db.Save(&request).Error; err != nil {
		return
	}

	response = request
	return
}

package asset_repo

import (
	"log"
	"test-kreditplus/src/entities"
	"time"

	"gorm.io/gorm"
)

type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) (*AssetRepository, error) {
	return &AssetRepository{
		db: db,
	}, nil
}

func (a AssetRepository) Create(db *gorm.DB, request *entities.Asset) (response *entities.Asset, err error) {

	if db != nil {
		a.db = db
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = a.db.Create(request).Error; err != nil {
		log.Printf("asset create error: %v", err)
		return
	}

	response = request
	return
}

func (a AssetRepository) CreateInBatches(db *gorm.DB, request []*entities.Asset) (response []*entities.Asset, err error) {

	if db != nil {
		a.db = db
	}

	for _, val := range request {
		val.CreatedAt = time.Now()
		val.UpdatedAt = time.Now()
	}

	if err = a.db.CreateInBatches(request, len(request)).Error; err != nil {
		log.Printf("asset create in batches error: %v", err)
		return
	}

	response = request

	return
}

func (a AssetRepository) Update(db *gorm.DB, request *entities.Asset) (response *entities.Asset, err error) {
	if db != nil {
		a.db = db
	}

	request.UpdatedAt = time.Now()
	if err = a.db.Save(request).Error; err != nil {
		log.Printf("asset Update error: %v", err)
		return
	}

	response = request
	return
}

func (a AssetRepository) Get(id uint) (response *entities.Asset, err error) {

	if err = a.db.Where("id = ?", id).First(response).Error; err != nil {
		log.Printf("asset get error: %v", err)
		return
	}

	return
}

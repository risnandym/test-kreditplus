package repositories

import (
	"kredit_plus/src/app/entities"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*UserRepository, error) {
	return &UserRepository{db: db}, nil
}

func (u UserRepository) Create(request entities.User) (response *entities.User, err error) {
	log.Println("berhasil masuk repo")

	// Convert UUID to string

	request.UUID = uuid.New()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	log.Println(u.db)
	response, err = request.SaveUser(u.db)
	if err != nil {
		return
	}
	log.Println("berhasil masuk db")
	return
}

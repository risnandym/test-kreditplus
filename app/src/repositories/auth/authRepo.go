package auth_repo

import (
	"kredit_plus/app/src/entities"
	"kredit_plus/core/config"
	"kredit_plus/core/utils"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db     *gorm.DB
	config config.Configuration
}

func NewAuthRepository(db *gorm.DB, config config.Configuration) (*AuthRepository, error) {
	return &AuthRepository{
		db:     db,
		config: config,
	}, nil
}

func (a AuthRepository) Create(request entities.Auth) (response *entities.Auth, err error) {

	request.UUID = uuid.New()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	response, err = request.SaveUser(a.db)
	if err != nil {
		return
	}

	return
}

func (a AuthRepository) Login(email string, password string) (token string, err error) {

	auth := entities.Auth{}

	err = a.db.Model(entities.Auth{}).Where("email = ?", email).Take(&auth).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err = utils.GenerateToken(a.config, auth.ID, auth.UUID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a AuthRepository) Get(id uint) (response entities.Auth, err error) {

	result := a.db.First(&response, "id = ?", id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return
}

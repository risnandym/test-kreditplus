package entities

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	IDForm
	UUID      uuid.UUID `gorm:"not null;unique" json:"uuid"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Phone     string    `gorm:"not null;" json:"phone"`
	Password  string    `gorm:"not null;" json:"password"`
	LastLogin time.Time `gorm:"not null;" json:"last_login"`
	TimeStamp

	Limit             Limit               `json:"-"`
	Profile           Profile             `json:"-"`
	CreditTransaction []CreditTransaction `json:"-"`
	Asset             []Asset             `json:"-"`
}

func (a *Auth) SaveUser(db *gorm.DB) (*Auth, error) {

	hashedPassword, errPassword := a.GeneratePassword(a.Password)
	if errPassword != nil {
		return &Auth{}, errPassword
	}

	a.Password = string(hashedPassword)

	var err error = db.Create(&a).Error
	if err != nil {
		return &Auth{}, err
	}

	return a, nil
}

func (a *Auth) GeneratePassword(pass string) (string, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if errPassword != nil {
		return "", errPassword
	}
	return string(hashedPassword), nil
}

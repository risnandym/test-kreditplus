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
}

func (u *Auth) SaveUser(db *gorm.DB) (*Auth, error) {

	hashedPassword, errPassword := u.GeneratePassword(u.Password)
	if errPassword != nil {
		return &Auth{}, errPassword
	}

	u.Password = string(hashedPassword)

	var err error = db.Create(&u).Error
	if err != nil {
		return &Auth{}, err
	}

	return u, nil
}

func (u *Auth) GeneratePassword(pass string) (string, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if errPassword != nil {
		return "", errPassword
	}
	return string(hashedPassword), nil
}

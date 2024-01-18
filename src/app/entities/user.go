package entities

import (
	"kredit_plus/core/utils/token"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UUID      uuid.UUID `gorm:"not null;unique" json:"uuid"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Phone     string    `gorm:"not null;" json:"phone"`
	Password  string    `gorm:"not null;" json:"password"`
	LastLogin time.Time `gorm:"not null;" json:"last_login"`
	TimeStamp
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := User{}

	err = db.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID, u.UUID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	//turn password into hash
	hashedPassword, errPassword := u.GeneratePassword(u.Password)
	if errPassword != nil {
		return &User{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	log.Println(u)

	var err error = db.Create(&u).Error
	if err != nil {

		return &User{}, err
	}
	return u, nil
}

func (u *User) GeneratePassword(pass string) (string, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if errPassword != nil {
		return "", errPassword
	}
	return string(hashedPassword), nil
}

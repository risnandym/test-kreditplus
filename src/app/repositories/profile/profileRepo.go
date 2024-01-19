package profile_repo

import (
	"test-kreditplus/src/app/entities"
	"time"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) (*ProfileRepository, error) {
	return &ProfileRepository{
		db: db,
	}, nil
}

func (p ProfileRepository) Create(request entities.Profile) (response *entities.Profile, err error) {

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	if err = p.db.Create(&request).Error; err != nil {
		return
	}

	return
}

// func (p ProfileRepository) Login(username string, password string) (token string, err error) {

// 	user := entities.Profile{}

// 	err = p.db.Model(entities.Profile{}).Where("email = ?", username).Take(&user).Error
// 	if err != nil {
// 		return "", err
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
// 		return "", err
// 	}

// 	token, err = utils.GenerateToken(p.config, user.ID, user.UUID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// }

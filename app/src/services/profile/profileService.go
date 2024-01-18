package profile_service

import (
	"kredit_plus/app/src/contract"
	"kredit_plus/app/src/entities"
)

func (u ProfileService) Create(request contract.ProfileInput) (response *contract.ProfileInput, err error) {

	profile := entities.Profile{}
	profile.AuthID = request.AuthID
	profile.NIK = request.NIK
	profile.FullName = request.FullName
	profile.LegalName = request.LegalName
	profile.PlaceOfBirth = request.PlaceOfBirth
	profile.DateOfBirth = request.DateOfBirth
	profile.Salary = request.Salary
	profile.KtpImage = request.KtpImage
	profile.SelfieImage = request.SelfieImage

	_, err = u.profileRepo.Create(profile)
	if err != nil {
		return nil, err
	}

	response = &request

	return
}

// func (u ProfileService) Login(request contract.LoginInput) (token string, err error) {

// 	profile := entities.Profile{}
// 	profile.Email = request.Email
// 	profile.Password = request.Password

// 	token, err = u.profileRepo.Login(profile.Email, profile.Password)

// 	return
// }

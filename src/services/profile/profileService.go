package profile_service

import (
	"test-kreditplus/src/contract"
	"test-kreditplus/src/entities"
)

func (p ProfileService) Create(request contract.ProfileInput) (response *contract.ProfileInput, err error) {

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

	_, err = p.profileRepo.Create(profile)
	if err != nil {
		return nil, err
	}

	var limit1, limit2, limit3, limit6 float32

	if profile.Salary*0.1 <= 750000 {
		limit1 = 100000
		limit2 = 200000
		limit3 = 500000
		limit6 = 700000
	} else {
		limit1 = 1000000
		limit2 = 1200000
		limit3 = 1500000
		limit6 = 2000000
	}

	limit := entities.Limit{
		AuthID: request.AuthID,
		Tenor1: limit1,
		Tenor2: limit2,
		Tenor3: limit3,
		Tenor4: limit6,
	}

	_, err = p.limitRepo.Create(limit)
	if err != nil {
		return nil, err
	}

	response = &request
	return
}

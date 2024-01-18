package app

import (
	auth_repo "kredit_plus/app/src/repositories/auth"
	profile_repo "kredit_plus/app/src/repositories/profile"
	auth_service "kredit_plus/app/src/services/auth"
	profile_service "kredit_plus/app/src/services/profile"
	"kredit_plus/core/config"
	"log"
)

type repositories struct {
	AuthRepo    *auth_repo.AuthRepository
	ProfileRepo *profile_repo.ProfileRepository
}

type services struct {
	AuthSVC    *auth_service.AuthService
	ProfileSVC *profile_service.ProfileService
}

type Dependency struct {
	Repositories *repositories
	Services     *services
}

func initRepositories() *repositories {
	var r repositories
	var err error

	r.AuthRepo, err = auth_repo.NewAuthRepository(config.DB(), config.Config())
	if err != nil {
		log.Panic(err)
	}

	r.ProfileRepo, err = profile_repo.NewProfileRepository(config.DB())
	if err != nil {
		log.Panic(err)
	}

	return &r
}

func initServices(r *repositories) *services {

	return &services{
		AuthSVC:    auth_service.NewAuthService(r.AuthRepo),
		ProfileSVC: profile_service.NewProfileService(r.ProfileRepo),
	}
}

func Dependencies() *Dependency {
	repositories := initRepositories()
	services := initServices(repositories)

	return &Dependency{
		Repositories: repositories,
		Services:     services,
	}
}

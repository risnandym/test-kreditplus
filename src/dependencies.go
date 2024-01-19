package app

import (
	"log"
	"test-kreditplus/core/config"
	auth_repo "test-kreditplus/src/app/repositories/auth"
	profile_repo "test-kreditplus/src/app/repositories/profile"
	auth_service "test-kreditplus/src/app/services/auth"
	profile_service "test-kreditplus/src/app/services/profile"
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

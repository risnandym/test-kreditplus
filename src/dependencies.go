package src

import (
	"kredit_plus/core/config"
	repo "kredit_plus/src/app/repositories"
	user_service "kredit_plus/src/app/services/user"
	"log"
)

type repositories struct {
	UserRepo *repo.UserRepository
}

type services struct {
	UserSVC *user_service.UserService
}

type Dependency struct {
	Repositories *repositories
	Services     *services
}

func initRepositories() *repositories {
	var r repositories
	var err error

	r.UserRepo, err = repo.NewUserRepository(config.DB())
	if err != nil {
		log.Panic(err)
	}

	return &r
}

func initServices(r *repositories) *services {
	// trxService := transactionSvc.TransactionServiceInput{}

	return &services{
		UserSVC: user_service.NewMerchantService(r.UserRepo),
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

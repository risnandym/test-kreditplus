package src

import (
	"log"
	"test-kreditplus/core/config"
	asset_repo "test-kreditplus/src/repositories/asset"
	auth_repo "test-kreditplus/src/repositories/auth"
	credit_repo "test-kreditplus/src/repositories/credit"
	debit_repo "test-kreditplus/src/repositories/debit"
	limit_repo "test-kreditplus/src/repositories/limit"
	profile_repo "test-kreditplus/src/repositories/profile"
	auth_service "test-kreditplus/src/services/auth"
	profile_service "test-kreditplus/src/services/profile"
	transaction_service "test-kreditplus/src/services/transaction"
)

type repositories struct {
	AuthRepo    *auth_repo.AuthRepository
	ProfileRepo *profile_repo.ProfileRepository
	LimitRepo   *limit_repo.LimitRepository
	CreditRepo  *credit_repo.CreditRepository
	DebitRepo   *debit_repo.DebitRepository
	AssetRepo   *asset_repo.AssetRepository
}

type services struct {
	AuthSVC        *auth_service.AuthService
	ProfileSVC     *profile_service.ProfileService
	TransactionSVC *transaction_service.TransactionService
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

	r.LimitRepo, err = limit_repo.NewLimitRepository(config.DB())
	if err != nil {
		log.Panic(err)
	}

	r.CreditRepo, err = credit_repo.NewCreditRepository(config.DB())
	if err != nil {
		log.Panic(err)
	}

	r.DebitRepo, err = debit_repo.NewDebitRepository(config.DB())
	if err != nil {
		log.Panic(err)
	}

	r.AssetRepo, err = asset_repo.NewAssetRepository(config.DB())
	if err != nil {
		log.Panic(err)
	}

	return &r
}

func initServices(r *repositories) *services {

	return &services{
		AuthSVC:        auth_service.NewAuthService(r.AuthRepo),
		ProfileSVC:     profile_service.NewProfileService(config.DB(), r.ProfileRepo, r.LimitRepo),
		TransactionSVC: transaction_service.NewTransactionService(config.DB(), r.CreditRepo, r.DebitRepo, r.LimitRepo, r.AssetRepo),
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

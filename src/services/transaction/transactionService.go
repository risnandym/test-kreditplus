package transaction_service

import (
	"log"
	"math"
	"slices"
	"test-kreditplus/src/contract"
	"test-kreditplus/src/entities"
	"time"

	"gorm.io/gorm"
)

func (t TransactionService) Credit(request contract.CreditInput) (response *contract.CreditOutput, err error) {

	limit, err := t.limitRepo.Get(request.AuthID)
	if err != nil {
		log.Printf("create transaction service err: %v", err)
		return nil, err
	}

	adminFee := request.OTRAmount * contract.AdminFee
	totalInstallment := adminFee + request.OTRAmount
	monthlyPayment := totalInstallment * (contract.InterestRate / (1 - 1/(math.Pow(1+contract.InterestRate, float64(request.InstallmentPeriod)))))
	totalRepayment := monthlyPayment * float64(request.InstallmentPeriod)
	totalInterest := totalRepayment - request.OTRAmount
	interest := totalInterest / float64(request.InstallmentPeriod)
	contractnumber := contract.GenerateContractID(request.AuthID)

	limit, err = contract.CalculateLimit(limit, totalInstallment, request.InstallmentPeriod)
	if err != nil {
		log.Printf("create transaction service err: %v", err)
		return nil, err
	}

	log.Print(limit)

	credit := &entities.Credit{
		AuthID:            request.AuthID,
		ContractNumber:    contractnumber,
		OTRAmount:         request.OTRAmount,
		AdminFee:          adminFee,
		InstallmentAmount: totalRepayment,
		InstallmentPeriod: request.InstallmentPeriod,
		Interest:          interest,
		TotalInterest:     totalInterest,
		SalesChannel:      request.SalesChannel,
	}

	asset := &entities.Asset{
		AuthID:         request.AuthID,
		ContractNumber: contractnumber,
	}

	debit := entities.Debit{
		AuthID:          request.AuthID,
		ContractNumbers: contractnumber,
		FinalAmount:     monthlyPayment,
		AmoutFixed:      false,
		Paid:            false,
	}

	err = t.db.Transaction(func(db *gorm.DB) error {

		var debits []entities.Debit
		var assets []*entities.Asset

		credit, err = t.creditRepo.Create(db, credit)
		if err != nil {
			log.Printf("create transaction service err: %v", err)
			return err
		}

		limit, err = t.limitRepo.Update(db, limit)
		if err != nil {
			log.Printf("create transaction service err: %v", err)
			return err
		}

		for _, val := range request.Assets {

			asset.Name = val.Name
			asset.Type = val.Type
			asset.Price = val.Price
			asset.Description = val.Description

			assets = append(assets, asset)
		}

		for i := 0; i < request.InstallmentPeriod; i++ {
			timenow := time.Now().AddDate(0, i, 0)
			month := timenow.Month().String()
			year := timenow.Year()
			debit.Month = month
			debit.Year = year
			err = t.UpdateOrCreateDebit(db, &debit)
			if err != nil {
				if err != gorm.ErrRecordNotFound {
					log.Printf("create transaction service err: %v", err)
					return err
				}
				debits = append(debits, debit)
			}
		}

		firstbatch := slices.Clone(debits)
		_, err = t.debitRepo.CreateInBatches(db, firstbatch, firstbatch[0].Month)
		if err != nil {
			log.Printf("create transaction service err: %v", err)
			return err
		}

		_, err = t.assetRepo.CreateInBatches(db, assets)
		if err != nil {
			log.Printf("create transaction service err: %v", err)
			return err
		}

		return err
	})
	if err != nil {
		log.Printf("create transaction service err: %v", err)
		return nil, err
	}

	response = &contract.CreditOutput{
		AuthID:            request.AuthID,
		ContractNumber:    credit.ContractNumber,
		OTRAmount:         credit.OTRAmount,
		AdminFee:          credit.AdminFee,
		InstallmentAmount: credit.InstallmentAmount,
		InstallmentPeriod: credit.InstallmentPeriod,
		Interest:          credit.Interest,
		TotalInterest:     credit.TotalInterest,
		SalesChannel:      credit.SalesChannel,
		Limit: contract.Limit{
			Tenor1: limit.Tenor1,
			Tenor2: limit.Tenor2,
			Tenor3: limit.Tenor3,
			Tenor6: limit.Tenor6,
		},
		Assets: request.Assets,
	}

	return
}

// func (t TransactionService) Debit(request contract.CreditInput) (response *contract.CreditOutput, err error) {

// 	limit, err := t.limitRepo.Get(uint(request.AuthID))
// 	if err != nil {
// 		log.Printf("create transaction service err: %v", err)
// 		return nil, err
// 	}

// 	// transaction := &entities.CreditTransaction{
// 	// 	UUID:              request.UUID,
// 	// 	AuthID:            request.AuthID,
// 	// 	AssetID:           request.AssetID,
// 	// 	ContractNumber:    request.ContractNumber,
// 	// 	OTRAmount:         request.AdminFee,
// 	// 	AdminFee:          request.AdminFee,
// 	// 	InstallmentAmount: request.InstallmentAmount,
// 	// 	InstallmentPeriod: request.InstallmentPeriod,
// 	// 	Interest:          request.InterestAmount,
// 	// 	SalesChannel:      request.SalesChannel,
// 	// }

// 	totalInstallment := request.InstallmentAmount + request.InterestAmount + request.AdminFee + request.OTRAmount

// 	limit, err = limit.CalculateLimit(totalInstallment, request.InstallmentPeriod)
// 	if err != nil {
// 		log.Printf("create transaction service err: %v", err)
// 		return nil, err
// 	}

// 	err = t.db.Transaction(func(db *gorm.DB) error {

// 		transaction, err = t.creditRepo.Create(db, transaction)
// 		if err != nil {
// 			log.Printf("create transaction service err: %v", err)
// 			return err
// 		}

// 		limit, err = t.limitRepo.Update(db, limit)
// 		if err != nil {
// 			log.Printf("create transaction service err: %v", err)
// 			return err
// 		}

// 		return err
// 	})
// 	if err != nil {
// 		log.Printf("create transaction service err: %v", err)
// 		return nil, err
// 	}

// 	response = &contract.CreditOutput{
// 		CreditInput: request,
// 		Limit:       *limit,
// 	}

// 	return
// }

func (t TransactionService) UpdateOrCreateDebit(db *gorm.DB, incoming *entities.Debit) error {
	current, err := t.debitRepo.Get(incoming.Month, incoming.Year)
	if err != nil {
		return err
	}

	current.ContractNumbers += "," + incoming.ContractNumbers
	current.FinalAmount += incoming.FinalAmount

	_, err = t.debitRepo.Update(db, current)

	return err
}

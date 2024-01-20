package transaction_service

import (
	"log"
	"test-kreditplus/src/contract"
	"test-kreditplus/src/entities"

	"gorm.io/gorm"
)

func (t TransactionService) Credit(request contract.TransactionInput) (response *contract.TransactionOutput, err error) {

	limit, err := t.limitRepo.Get(uint(request.AuthID))
	if err != nil {
		log.Printf("create transaction service err: %v", err)
		return nil, err
	}

	transaction := &entities.CreditTransaction{
		UUID:              request.UUID,
		AuthID:            request.AuthID,
		AssetID:           request.AssetID,
		ContractNumber:    request.ContractNumber,
		OTRAmount:         request.AdminFee,
		AdminFee:          request.AdminFee,
		InstallmentAmount: request.InstallmentAmount,
		InstallmentPeriod: request.InstallmentPeriod,
		InterestAmount:    request.InterestAmount,
		SalesChannel:      request.SalesChannel,
	}

	totalInstallment := request.InstallmentAmount + request.InterestAmount + request.AdminFee + request.OTRAmount

	limit, err = limit.CalsulateLimit(totalInstallment, request.InstallmentPeriod)
	if err != nil {
		log.Printf("create transaction service err: %v", err)
		return nil, err
	}

	err = t.db.Transaction(func(db *gorm.DB) error {

		transaction, err = t.creditRepo.Create(db, transaction)
		if err != nil {
			log.Printf("create transaction service err: %v", err)
			return err
		}

		limit, err = t.limitRepo.Update(db, limit)
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

	response = &contract.TransactionOutput{
		TransactionInput: request,
		Limit:            *limit,
	}

	return
}

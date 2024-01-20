package contract

import (
	"test-kreditplus/src/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionInput struct {
	UUID              uuid.UUID `json:"uuid"`
	AuthID            int       `json:"auth_id"`
	AssetID           *int      `json:"asset_id"`
	ContractNumber    string    `json:"contract_number"`
	OTRAmount         float32   `json:"otr_amount"`
	AdminFee          float32   `json:"admin_fee"`
	InstallmentAmount float32   `json:"installment_amount"`
	InstallmentPeriod int       `json:"installment_period"`
	InterestAmount    float32   `json:"interest_amount"`
	SalesChannel      string    `json:"sales_channel"`
}

type TransactionOutput struct {
	TransactionInput
	Limit entities.Limit `json:"limit"`
}

func ValidateAndBuildCreditInput(c *gin.Context) (request TransactionInput, err error) {
	authctx, _ := c.Get("auth")
	auth := authctx.(entities.Auth)

	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	request.AuthID = int(auth.ID)
	return
}

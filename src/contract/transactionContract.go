package contract

import (
	"fmt"
	"strconv"
	"strings"
	"test-kreditplus/src/entities"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	InterestRate float64 = 0.03
	AdminFee     float64 = 0.01
)

type CreditInput struct {
	AuthID            uint    `json:"auth_id" validate:"required"`
	OTRAmount         float64 `json:"otr_amount" validate:"required"`
	InstallmentPeriod int     `json:"installment_period" validate:"required"`
	SalesChannel      string  `json:"sales_channel" validate:"required"`
	Assets            []Asset `json:"assets" validate:"required"`
}

type CreditOutput struct {
	AuthID            uint    `json:"auth_id"`
	ContractNumber    string  `json:"contract_number"`
	OTRAmount         float64 `json:"otr_amount"`
	AdminFee          float64 `json:"admin_fee"`
	InstallmentAmount float64 `json:"installment_amount"`
	InstallmentPeriod int     `json:"installment_period"`
	Interest          float64 `json:"interest"`
	TotalInterest     float64 `json:"total_interest"`
	SalesChannel      string  `json:"sales_channel"`
	Limit             Limit   `json:"limit"`
	Assets            []Asset `json:"assets"`
}

func ValidateAndBuildCreditInput(c *gin.Context) (request CreditInput, err error) {
	authctx, _ := c.Get("auth")
	auth := authctx.(entities.Auth)

	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	request.AuthID = auth.ID
	return
}

func ValidateAndBuildDebitInput(c *gin.Context) (request CreditInput, err error) {
	authctx, _ := c.Get("auth")
	auth := authctx.(entities.Auth)

	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	request.AuthID = auth.ID
	return
}

func GenerateContractID(authID uint) string {

	head := "CONTRACT"
	container, _ := strconv.Atoi(fmt.Sprintf("%d%d", authID, time.Now().UnixMilli()))
	head += fmt.Sprintf("%012v", strconv.FormatInt(int64(container), 36))
	head = strings.ToUpper(head)

	return head
}

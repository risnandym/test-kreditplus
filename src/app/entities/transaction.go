package entities

import "github.com/google/uuid"

type Transaction struct {
	IDForm
	UUID              uuid.UUID `json:"uuid"`
	CustomerID        int       `json:"customer_id"`
	AssetID           *int      `json:"asset_id"`
	ContractNumber    string    `json:"contract_number"`
	OTRAmount         float32   `json:"otr_amount"`
	AdminFee          float32   `json:"admin_fee"`
	InstallmentAmount float32   `json:"installment_amount"`
	InstallmentPeriod int       `json:"installment_period"`
	InterestAmount    float32   `json:"interest_amount"`
	SalesChannel      string    `json:"sales_channel"`
	TimeStamp
}

package entities

type Credit struct {
	IDForm
	AuthID            uint    `json:"auth_id"`
	ContractNumber    string  `json:"contract_number"`
	OTRAmount         float64 `json:"otr_amount"`
	AdminFee          float64 `json:"admin_fee"`
	InstallmentAmount float64 `json:"installment_amount"`
	InstallmentPeriod int     `json:"installment_period"`
	Interest          float64 `json:"interest"`
	TotalInterest     float64 `json:"total_interest"`
	SalesChannel      string  `json:"sales_channel"`
	TimeStamp
}

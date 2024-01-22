package entities

type Debit struct {
	IDForm
	AuthID          uint    `json:"auth_id"`
	ContractNumbers string  `json:"contract_numbers"`
	FinalAmount     float64 `json:"final_amount"`
	Month           string  `json:"month"`
	Year            int     `json:"year"`
	AmoutFixed      bool    `json:"amout_fixed"`
	Paid            bool    `json:"paid"`
	TimeStamp
}

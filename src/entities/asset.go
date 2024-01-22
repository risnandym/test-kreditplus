package entities

type Asset struct {
	IDForm
	AuthID         uint    `json:"auth_id"`
	ContractNumber string  `json:"contract_number"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	Paid           bool    `json:"paid"`
	TimeStamp
}

package entities

type UserLimit struct {
	IDForm
	CustomerID  int     `json:"customer_id"`
	Tenor       int     `json:"tenor"`
	LimitAmount float32 `json:"limit_amount"`
	TimeStamp
}

package entities

type Limit struct {
	IDForm
	AuthID      int     `json:"auth_id"`
	Tenor       int     `json:"tenor"`
	LimitAmount float32 `json:"limit_amount"`
	TimeStamp

	Auth Auth `json:"-"`
}

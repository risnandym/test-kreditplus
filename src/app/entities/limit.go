package entities

type Limit struct {
	IDForm
	AuthID int     `json:"auth_id"`
	Tenor1 float32 `json:"tenor1"`
	Tenor2 float32 `json:"tenor2"`
	Tenor3 float32 `json:"tenor3"`
	Tenor4 float32 `json:"tenor4"`
	TimeStamp

	Auth Auth `json:"-"`
}

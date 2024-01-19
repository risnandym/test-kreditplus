package entities

type Asset struct {
	IDForm
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	TimeStamp
}

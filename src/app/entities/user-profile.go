package entities

type UserProfile struct {
	IDForm
	CustomerID   int     `json:"customer_id"`
	NIK          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float32 `json:"salary"`
	KtpImage     string  `json:"ktp_image"`
	SelfieImage  string  `json:"selfie_image"`
	TimeStamp
}

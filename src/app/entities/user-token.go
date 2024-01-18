package entities

import "time"

type UserToken struct {
	IDForm
	CustomerID            int       `json:"customer_id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	UserAgent             string    `json:"user_agent"`
	IPAddress             string    `json:"ip_address"`
	AccessTokenExpiredAt  time.Time `json:"access_token_expired_at"`
	RefreshTokenExpiredAt time.Time `json:"refresh_token_expired_at"`
	TimeStamp
}

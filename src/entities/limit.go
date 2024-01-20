package entities

import (
	"errors"
	"log"
)

type Limit struct {
	IDForm
	AuthID int     `json:"auth_id"`
	Tenor1 float32 `json:"tenor1"`
	Tenor2 float32 `json:"tenor2"`
	Tenor3 float32 `json:"tenor3"`
	Tenor6 float32 `json:"tenor4"`
	TimeStamp
}

func (l *Limit) CalsulateLimit(amount float32, tenor int) (*Limit, error) {

	var err error
	var limitOnInstallment float32

	switch tenor {
	case 1:
		limitOnInstallment = l.Tenor1
	case 2:
		limitOnInstallment = l.Tenor2
	case 3:
		limitOnInstallment = l.Tenor3
	case 6:
		limitOnInstallment = l.Tenor6
	}

	if amount > limitOnInstallment {
		err = errors.New("transaksi melebihi limit pinjaman")
		log.Println(err)
		return l, err
	}

	l.Tenor1 -= limitOnInstallment
	l.Tenor2 -= limitOnInstallment
	l.Tenor3 -= limitOnInstallment
	l.Tenor6 -= limitOnInstallment

	if l.Tenor1 < 0 {
		l.Tenor1 = 0
	}

	if l.Tenor2 < 0 {
		l.Tenor2 = 0
	}

	if l.Tenor3 < 0 {
		l.Tenor3 = 0
	}

	return l, nil
}

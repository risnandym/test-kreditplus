package entities

import (
	"errors"
	"log"
)

type Limit struct {
	IDForm
	AuthID uint    `json:"auth_id"`
	Tenor1 float64 `json:"tenor1"`
	Tenor2 float64 `json:"tenor2"`
	Tenor3 float64 `json:"tenor3"`
	Tenor6 float64 `json:"tenor4"`
	TimeStamp
}

func CalculateLimit(l *Limit, amount float64, tenor int) (*Limit, error) {

	var err error
	var limitOnInstallment float64

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

	log.Print(l)

	if amount > limitOnInstallment {
		err = errors.New("transaksi melebihi limit pinjaman")
		log.Println(err)
		return l, err
	}

	l.Tenor1 -= limitOnInstallment
	log.Print(l.Tenor1, limitOnInstallment)
	l.Tenor2 -= limitOnInstallment
	log.Print(l.Tenor2, limitOnInstallment)
	l.Tenor3 -= limitOnInstallment
	log.Print(l.Tenor3, limitOnInstallment)
	l.Tenor6 -= limitOnInstallment
	log.Print(l.Tenor6, limitOnInstallment)

	log.Print(l)

	if l.Tenor1 < 0 {
		l.Tenor1 = 0
	}

	if l.Tenor2 < 0 {
		l.Tenor2 = 0
	}

	if l.Tenor3 < 0 {
		l.Tenor3 = 0
	}

	log.Print(l)
	log.Print(3)

	return l, nil
}

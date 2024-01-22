package contract

import (
	"errors"
	"log"
	"test-kreditplus/src/entities"
)

type Limit struct {
	Tenor1 float64 `json:"tenor1"`
	Tenor2 float64 `json:"tenor2"`
	Tenor3 float64 `json:"tenor3"`
	Tenor6 float64 `json:"tenor6"`
}

func CalculateLimit(l *entities.Limit, amount float64, tenor int) (*entities.Limit, error) {

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

	if amount > limitOnInstallment {
		err = errors.New("transaksi melebihi limit pinjaman")
		log.Println(err)
		return l, err
	}

	l.Tenor1 -= amount
	l.Tenor2 -= amount
	l.Tenor3 -= amount
	l.Tenor6 -= amount

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
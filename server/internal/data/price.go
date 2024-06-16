package data

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Price decimal.Decimal

var ErrorInvalidPrice = errors.New("invalid price")

func (p *Price) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "" {
		return ErrorInvalidPrice
	}

	dec, err := decimal.NewFromString(string(data))
	if err != nil {
		return err
	}

	*p = Price(dec.Round(4))
	return nil
}

func (p *Price) MarshalJSON() ([]byte, error) {
	strValue := decimal.Decimal(*p).StringFixed(4)
	return []byte(strValue), nil
}

func (p *Price) Value() string {
	return decimal.Decimal(*p).StringFixed(4)
}

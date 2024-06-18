package data

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Price int64

var ErrorInvalidPrice = errors.New("must be granter than or equal to 0")

func (p *Price) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "" {
		return ErrorInvalidPrice
	}

	dec, err := decimal.NewFromString(string(data))
	if err != nil {
		return err
	}
	intValue := dec.Round(4).Mul(decimal.NewFromInt(10000)).IntPart()

	*p = Price(intValue)
	return nil
}

func (p *Price) MarshalJSON() ([]byte, error) {
	dec := decimal.NewFromInt(int64(*p)).Div(decimal.NewFromInt(10000)).StringFixed(4)
	return []byte(dec), nil
}

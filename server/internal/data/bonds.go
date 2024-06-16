package data

import (
	"context"
	"database/sql"
	"time"

	"boundsApp.victorinolavida/internal/validator"
	"github.com/shopspring/decimal"
)

type Bond struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       Price  `json:"price"`
	NumberBonds int    `json:"number_bonds"`
}
type BondModel struct {
	DB *sql.DB
}

func ValidateBond(v *validator.Validator, bond *Bond) {
	v.Check(bond.Name != "", "name", "must be provided")
	v.Check(len(bond.Name) <= 40, "name", "must not be more than 40 characters long")
	v.Check(len(bond.Name) >= 3, "name", "must be at least 3 characters long")
	validatePrice(v, bond.Price)
	v.Check(bond.NumberBonds > 0, "number_bonds", "must be greater than zero")
	v.Check(bond.NumberBonds <= 10_000, "number_bonds", "must be less than or equal to 10000")
}

func validatePrice(v *validator.Validator, price Price) {
	minValue := decimal.NewFromInt(0)
	maxValue := decimal.NewFromInt(100_000_000)
	v.Check(minValue.Compare(decimal.Decimal(price)) == -1, "price", "must be greater than zero")
	v.Check(maxValue.Compare(decimal.Decimal(price)) == 1, "price", "must be less 100,000,000")
}

func (m *BondModel) Insert(bond *Bond, user *User) (*Bond, error) {
	query := `INSERT INTO 
	bonds (name, price, number_bonds, issuer) 
	VALUES ($1, $2, $3, $4) RETURNING id`
	args := []any{bond.Name, bond.Price.Value(), bond.NumberBonds, user.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&bond.ID)

	if err != nil {
		return nil, err
	}
	return bond, nil
}

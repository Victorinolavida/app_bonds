package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"boundsApp.victorinolavida/internal/validator"
	"github.com/shopspring/decimal"
)

var (
	ErrBoughtAlreadyBought = errors.New("Bond already bought or not available")
)

type Bond struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       Price     `json:"price"`
	NumberBonds int       `json:"number_bonds"`
	CreatedAt   time.Time `json:"created_at"`
	IssuerID    int64     `json:"-"`
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
	bonds (name, price, number_bonds, issuer_id) 
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
func (m *BondModel) GetBondsByUser(user User) ([]*Bond, error) {
	query := `
		SELECT bonds.id, name, price, number_bonds, created_at FROM bonds
		LEFT JOIN transactions ON bonds.id = transactions.bond_id
		WHERE (issuer_id = $1 AND transactions.id is null) 
		OR (transactions.buyer_id= $1 AND transactions.deleted_at IS NULL)
		
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, user.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return []*Bond{}, nil
		default:
			return nil, err
		}
	}

	defer rows.Close()

	bonds := []*Bond{}

	for rows.Next() {
		var bond Bond
		var price float64
		err := rows.Scan(&bond.ID, &bond.Name, &price, &bond.NumberBonds, &bond.CreatedAt)

		decimalPrice := decimal.NewFromFloat(float64(price))
		bond.Price = Price(decimalPrice)

		if err != nil {
			return nil, err
		}

		bonds = append(bonds, &bond)
	}

	return bonds, nil

}

func (m *BondModel) GetBondByID(bond *Bond) error {
	query := `
	SELECT id, name, price, number_bonds, created_at, issuer_id FROM bonds
	WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var price float64
	err := m.DB.QueryRowContext(ctx, query, bond.ID).Scan(&bond.ID, &bond.Name, &price, &bond.NumberBonds, &bond.CreatedAt, &bond.IssuerID)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}

	decimalPrice := decimal.NewFromFloat(float64(price))
	bond.Price = Price(decimalPrice)

	return nil
}

func (m *BondModel) IsPurchasableBound(bond *Bond, user *User, transaction *Transaction) error {

	query := `
	SELECT transactions.buyer_id FROM bonds
		LEFT JOIN transactions ON bonds.id = transactions.bond_id
		WHERE (
		(issuer_id <> $1 AND transactions.id is null) 
		OR (transactions.buyer_id <> $1 AND transactions.deleted_at IS NULL)
		) AND bonds.id= $2
		
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var price float64
	err := m.DB.QueryRowContext(ctx, query, user.ID, bond.ID).Scan(&transaction.SellerId)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrBoughtAlreadyBought
		default:
			return err
		}
	}

	decimalPrice := decimal.NewFromFloat(float64(price))
	bond.Price = Price(decimalPrice)

	return nil

}

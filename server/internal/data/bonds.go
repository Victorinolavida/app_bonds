package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"boundsApp.victorinolavida/internal/validator"
)

var (
	ErrBoughtAlreadyBought = errors.New("bond already bought or not available")
)

type Bond struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       Price     `json:"price"`
	NumberBonds int       `json:"number_bonds"`
	OwnerId     int64     `json:"-"`
	CreatedBy   int64     `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
}
type BondWithOwner struct {
	Bond
	Owner string `json:"owner"`
}
type BondWithStatus struct {
	Bond
	Status string `json:"status"`
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
	v.Check(bond.NumberBonds <= 10_000, "number_bonds", "must be less than or equal to 10,000")
}

func validatePrice(v *validator.Validator, price Price) {
	v.Check(price <= 100_000_000*10_000, "price", "must be less than or equal to 100,000,000")
	v.Check(price >= 0, "price", "must be greater than or equal to 0")
}
func (m *BondModel) Insert(bond *Bond) error {
	query := `INSERT INTO 
	bonds (name, price, number_bonds, owner_id, created_by) 
	VALUES ($1, $2, $3, $4, $4) 
	RETURNING id,created_at`
	args := []any{bond.Name, bond.Price, bond.NumberBonds, bond.OwnerId}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&bond.ID, &bond.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
func (m *BondModel) ChangeOwner(bond *Bond, user *User) error {
	query := `UPDATE bonds SET owner_id = $1 WHERE id = $2`
	args := []any{user.ID, bond.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (m *BondModel) GetBondsByUser(user User, pagination Pagination) ([]*BondWithStatus, Pagination, error) {
	query := `
	SELECT COUNT(*) OVER(), id, name, price, number_bonds, owner_id, created_at, 
		case when owner_id = created_by then 'CREATED' else 'BOUGHT' end as role
	FROM bonds 
	WHERE owner_id = $1 
	ORDER BY created_at DESC
	LIMIT $2 OFFSET $3
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, user.ID, pagination.limit(), pagination.offset())
	if err != nil {
		return nil, Pagination{}, err
	}

	defer rows.Close()

	totalRecords := 0
	bonds := []*BondWithStatus{}

	for rows.Next() {
		var bond BondWithStatus
		err := rows.Scan(&totalRecords, &bond.ID, &bond.Name, &bond.Price, &bond.NumberBonds, &bond.OwnerId, &bond.CreatedAt, &bond.Status)

		if err != nil {
			return nil, Pagination{}, err
		}

		bonds = append(bonds, &bond)
	}

	paginationData := getPagination(totalRecords, pagination.CurrentPage, pagination.PageSize)
	return bonds, paginationData, nil

}

func (m *BondModel) GetBondByID(bond *Bond) error {
	query := `
	SELECT id, name, price, number_bonds, created_at, owner_id FROM bonds
	WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, bond.ID).Scan(&bond.ID, &bond.Name, &bond.Price, &bond.NumberBonds, &bond.CreatedAt, &bond.OwnerId)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}

	return nil
}

func (m *BondModel) GetAllAvailable(user *User, pagination Pagination) ([]*BondWithOwner, Pagination, error) {
	query := `
	SELECT COUNT(*) OVER(), bonds.id, name, price, number_bonds, bonds.created_at, users.username as owner  FROM bonds
	INNER JOIN users on bonds.owner_id = users.id
	WHERE owner_id <> $1
	ORDER BY created_at DESC
	LIMIT $2 OFFSET $3
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{user.ID, pagination.limit(), pagination.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Pagination{}, err
	}

	defer rows.Close()

	bonds := []*BondWithOwner{}
	totalRecords := 0

	for rows.Next() {
		var bond BondWithOwner
		var price float64
		err := rows.Scan(&totalRecords, &bond.ID, &bond.Name, &price, &bond.NumberBonds, &bond.CreatedAt, &bond.Owner)

		if err != nil {
			return nil, Pagination{}, err
		}

		bonds = append(bonds, &bond)
	}

	paginationData := getPagination(totalRecords, pagination.CurrentPage, pagination.PageSize)
	return bonds, paginationData, nil
}

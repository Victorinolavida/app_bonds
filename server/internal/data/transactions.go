package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Transaction struct {
	ID       int64     `json:"id"`
	BondId   string    `json:"bond_id"`
	SellerId int64     `json:"seller_id"`
	BuyerID  int64     `json:"buyer_id"`
	DeleteAt time.Time `json:"delete_at"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (m *TransactionModel) Insert(transaction *Transaction) (*Transaction, error) {

	query := `INSERT INTO 
	transactions (bond_id, seller_id, buyer_id) 
	VALUES ($1, $2, $3) RETURNING id`

	args := []any{transaction.BondId, transaction.SellerId, transaction.BuyerID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&transaction.ID)

	if err != nil {
		return nil, err
	}

	// set deleted_at= now() form previous transactions with the same bond_id
	query = `UPDATE transactions 
	SET deleted_at = now() 
	WHERE bond_id = $1 
	AND deleted_at IS NULL AND id <> $2`

	_, err = m.DB.ExecContext(ctx, query, transaction.BondId, transaction.ID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionModel) GetByBondId(bond *Bond, buyer *User) error {
	query := `SELECT id, bond_id, seller_id, buyer_id, deleted_at
	FROM transactions
	WHERE bond_id = $1 AND buyer_id = $2 AND deleted_at is NULL`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := t.DB.ExecContext(ctx, query, bond.ID, buyer.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	fmt.Printf("rows affected: %d\n", rowsAffected)
	if err != nil {
		return err
	}

	if rowsAffected >= 1 {
		return ErrBoughtAlreadyBought
	}

	return nil

}

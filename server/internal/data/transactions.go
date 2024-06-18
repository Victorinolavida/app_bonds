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
	Status   string    `json:"status"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (m *TransactionModel) Insert(transaction *Transaction) error {

	query := `INSERT INTO
	transactions (bond_id, seller_id, buyer_id, status)
	VALUES ($1, $2, $3, $4) RETURNING id,status`

	args := []any{transaction.BondId, transaction.SellerId, transaction.BuyerID, transaction.Status}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&transaction.ID, &transaction.Status)

	if err != nil {
		return err
	}

	return nil
}

func (m *TransactionModel) DeleteTransactions(transaction *Transaction) error {
	query := `UPDATE transactions SET deleted_at = NOW(),
	status = 'DELETED'
	WHERE bond_id = $1 AND buyer_id<>$2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, transaction.BondId, transaction.BuyerID)
	if err != nil {
		return err
	}
	return nil

}

func (m *TransactionModel) IsAlreadyBought(bond *Bond, buyer *User) error {

	query := `SELECT * from transactions 
	WHERE bond_id = $1 AND buyer_id = $2 AND deleted_at IS NULL`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, bond.ID, buyer.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected >= 1 {
		return ErrBoughtAlreadyBought
	}

	return nil

}

func (t *TransactionModel) GetByBondId(bond *Bond, buyer *User) error {
	query := `SELECT bonds.id, bond_id, seller_id, buyer_id, deleted_at
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

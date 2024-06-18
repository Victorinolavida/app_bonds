package data

import (
	"context"
	"database/sql"
	"time"
)

type BondTransaction struct {
	ID            int64  `json:"id"`
	BondID        string `json:"bond_id"`
	TransactionID int64  `json:"transaction_id"`
}

type BondTransactionModel struct {
	DB *sql.DB
}

func (m *BondTransactionModel) Insert(bondTransaction *BondTransaction) error {
	query := `INSERT INTO bond_transaction (bond_id, transaction_id) VALUES ($1, $2) RETURNING id`
	args := []any{bondTransaction.BondID, bondTransaction.TransactionID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&bondTransaction.ID)
}

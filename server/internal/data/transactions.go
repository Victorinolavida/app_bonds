package data

import (
	"context"
	"database/sql"
	"time"
)

type Transaction struct {
	ID       int64 `json:"id"`
	Price    Price `json:"price"`
	SellerID int64 `json:"seller_id"`
	BuyerID  int64 `json:"buyer_id"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (m *TransactionModel) Insert(transaction *Transaction) error {
	query := `INSERT INTO transactions (price, seller_id, buyer_id) VALUES ($1, $2, $3) RETURNING id`
	args := []any{transaction.Price, transaction.SellerID, transaction.BuyerID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&transaction.ID)
}

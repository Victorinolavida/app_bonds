package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("the requested record was not found")
)

type Models struct {
	Users           UserModel
	Bonds           BondModel
	Transactions    TransactionModel
	BondTransaction BondTransactionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:           UserModel{DB: db},
		Bonds:           BondModel{DB: db},
		Transactions:    TransactionModel{DB: db},
		BondTransaction: BondTransactionModel{DB: db},
	}
}

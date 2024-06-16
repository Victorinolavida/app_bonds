package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Users UserModel
	Bonds BondModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users: UserModel{DB: db},
		Bonds: BondModel{DB: db},
	}
}

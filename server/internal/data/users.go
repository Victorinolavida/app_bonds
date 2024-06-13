package data

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	CreatedAt    time.Time `json:"created_at"`
	PasswordHash []byte    `json:"-"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert() {

}

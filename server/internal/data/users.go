package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"boundsApp.victorinolavida/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

var ErrDuplicateEmail = errors.New("duplicate email")
var ErrDuplicateUsername = errors.New("duplicate username")

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Password  password  `json:"-"`
}

type password struct {
	plainText *string
	hash      []byte
}

type UserModel struct {
	DB *sql.DB
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Username != "", "username", "must be provided")
	v.Check(len(user.Username) >= 4, "username", "must be at least 8 characters long")
	v.Check(len(user.Username) <= 50, "username", "must not be more than 50 characters long")

	ValidateEmail(v, user.Email)

	if user.Password.plainText != nil {
		ValidatePassword(v, *user.Password.plainText)
	}
}
func ValidatePassword(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 characters long")
	v.Check(len(password) <= 50, "password", "must not be more than 50 characters long")
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(
		validator.Matches(email, *validator.EmailRex), "email", "must be a valid email address",
	)
}

func (p *password) SetPassword(plainText string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), 2)
	if err != nil {
		return err
	}

	p.plainText = &plainText
	p.hash = hash
	return nil
}

func (p *password) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func (m *UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users(
			email, username, password_hash
		)
		VALUES($1, $2, $3)
		RETURNING id, created_at
	`

	args := []any{user.Email, user.Username, user.Password.hash}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		switch {
		case err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"":
			return ErrDuplicateEmail
		case err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"":
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (m *UserModel) GetByEmail(email string) (*User, error) {
	var user User
	query := `
		SELECT id, email, created_at, password_hash, username
		FROM users 
		where email = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.CreatedAt,
		&user.Password.hash, &user.Username)
	if err != nil {

		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m *UserModel) GetByUsername(username string) (*User, error) {
	var user User
	query := `
		SELECT id, email, created_at, password_hash, username
		FROM users 
		where username = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Email, &user.CreatedAt,
		&user.Password.hash, &user.Username)
	if err != nil {

		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

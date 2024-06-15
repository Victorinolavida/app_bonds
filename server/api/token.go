package main

import (
	"errors"
	"net/http"
	"time"

	"boundsApp.victorinolavida/internal/data"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type CustomClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (app *application) createJWT(user *data.User) (string, error) {
	claims := CustomClaims{
		user.ID,
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(app.config.secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (app *application) SetCookieSession(w http.ResponseWriter, token string) {
	expires := time.Now().Add(time.Hour * 1)

	cookie := http.Cookie{
		Name:     TokenName,
		Domain:   "localhost",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expires,
	}

	cookie.Value = token

	http.SetCookie(w, &cookie)
}

func (app *application) removeCookieSession(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     TokenName,
		Domain:   "localhost",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
	}

	cookie.Value = ""
	http.SetCookie(w, &cookie)
}

func (app *application) validateToken(tokenString string, claims *CustomClaims) error {

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.config.secret), nil
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		return ErrInvalidCredentials
	}

	return nil
}

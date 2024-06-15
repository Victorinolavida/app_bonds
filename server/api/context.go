package main

import (
	"context"
	"net/http"

	"boundsApp.victorinolavida/internal/data"
)

type contextKey string

const contextKeyUser = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), contextKeyUser, user)

	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(contextKeyUser).(*data.User)
	if !ok {
		return nil
	}
	return user
}

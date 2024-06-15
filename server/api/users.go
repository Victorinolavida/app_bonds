package main

import (
	"errors"
	"net/http"

	"boundsApp.victorinolavida/internal/data"
	"boundsApp.victorinolavida/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := app.ReadJson(w, r, &input)
	validator := validator.New()

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Email:    input.Email,
		Username: input.Username,
	}
	err = user.Password.SetPassword(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	data.ValidateUser(validator, user)

	if !validator.IsValid() {
		app.fieldValidationResponse(w, r, validator.Errors)
		return
	}

	err = app.models.Users.Insert(user)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			validator.AddError("email", "this email is already in use")
			app.fieldValidationResponse(w, r, validator.Errors)
			return
		case errors.Is(err, data.ErrDuplicateUsername):
			validator.AddError("username", "this username is already in use")
			app.fieldValidationResponse(w, r, validator.Errors)
			return

		default:
			app.serverErrorResponse(w, r, err)
			return
		}

	}
	token, err := app.createJWT(user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	userData := envelop{"id": user.ID, "username": user.Username}
	app.SetCookieSession(w, token)
	app.WriteJson(w, http.StatusOK, envelop{"token": token, "user": userData})
}
func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email         string `json:"email"`
		PlainPassword string `json:"password"`
	}
	err := app.ReadJson(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateEmail(v, input.Email)
	v.Check(input.PlainPassword != "", "password", "must be provided")

	if !v.IsValid() {
		app.fieldValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidCredentialsResponse(w, r)
			return
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}
	//validate password
	ok, err := user.Password.ComparePassword(input.PlainPassword)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if !ok {
		app.invalidCredentialsResponse(w, r)
		return
	}
	token, err := app.createJWT(user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	userData := envelop{"id": user.ID, "username": user.Username}
	app.SetCookieSession(w, token)
	app.WriteJson(w, http.StatusOK, envelop{"token": token, "user": userData})
}

func (app *application) logoutUserHandler(w http.ResponseWriter, r *http.Request) {
	app.removeCookieSession(w)
	app.WriteJson(w, http.StatusOK, nil)
}

func (app *application) validateTokenHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user == nil {
		app.invalidCredentialsResponse(w, r)
		return
	}
	newToken, err := app.createJWT(user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	app.SetCookieSession(w, newToken)
	app.WriteJson(w, http.StatusOK, envelop{"token": newToken, "user": user})
}

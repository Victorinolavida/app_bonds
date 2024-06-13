package main

import (
	"net/http"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	// data := app.ReadJson(w,r,)
	// fmt.Printf("%+v", data)
	w.Write([]byte("register user page"))
}
func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login page"))
}

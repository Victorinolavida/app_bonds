package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/healthcheck", app.healthcheck)

	// signUP
	router.HandlerFunc(http.MethodPost, "/api/auth/join", app.registerUserHandler)
	// login
	router.HandlerFunc(http.MethodPost, "/api/auth/login", app.loginUserHandler)
	//logout
	router.HandlerFunc(http.MethodPost, "/api/auth/logout", app.authenticate(app.logoutUserHandler))

	router.HandlerFunc(http.MethodGet, "/api/auth/session", app.authenticate(app.validateTokenHandler))

	// create a bound
	router.HandlerFunc(http.MethodPost, "/api/bound", app.createBondHandler)
	// buy a bound
	router.HandlerFunc(http.MethodPut, "/api/bound/:id/buy", app.buyABondById)
	// list bounds
	router.HandlerFunc(http.MethodGet, "/api/bounds", app.listAllBonds)

	router.HandlerFunc(http.MethodGet, "/api/users/:id/bounds", app.listAllBondsCreatedByUser)

	return app.enableCors(router)
}

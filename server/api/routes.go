package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/healthcheck", app.healthcheck)

	// signUP
	router.HandlerFunc(http.MethodPost, "/api/join", app.registerUserHandler)
	// login
	router.HandlerFunc(http.MethodPost, "/api/login", app.loginUserHandler)

	// create a bound
	router.HandlerFunc(http.MethodPost, "/api/bound", app.createBondHandler)
	// buy a bound
	router.HandlerFunc(http.MethodPut, "/api/bound/:id/buy", app.buyABondById)
	// list bounds
	router.HandlerFunc(http.MethodGet, "/api/bounds", app.listAllBonds)

	router.HandlerFunc(http.MethodGet, "/api/users/:id/bounds", app.listAllBondsCreatedByUser)

	return router
}

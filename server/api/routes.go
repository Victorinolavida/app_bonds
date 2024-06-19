package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.HandlerFunc(http.MethodGet, "/api/healthcheck", app.healthcheck)

	// signUP
	router.HandlerFunc(http.MethodPost, "/api/auth/join", app.registerUserHandler)
	// login
	router.HandlerFunc(http.MethodPost, "/api/auth/login", app.loginUserHandler)
	//logout
	router.HandlerFunc(http.MethodPost, "/api/auth/logout", app.authenticate(app.logoutUserHandler))

	router.HandlerFunc(http.MethodGet, "/api/auth/session", app.authenticate(app.validateTokenHandler))

	// create a bound
	router.HandlerFunc(http.MethodPost, "/api/bonds", app.authenticate(app.createBondHandler))
	// buy a bound
	router.HandlerFunc(http.MethodPut, "/api/bonds/:id/buy", app.authenticate(app.buyBondHandler))

	// // list bounds owned by user
	router.HandlerFunc(http.MethodGet, "/api/bonds", app.authenticate(app.listOwnedBondsLoggedUserHandler))

	// list all purchasable bounds
	router.HandlerFunc(http.MethodGet, "/api/bonds/purchasable", app.authenticate(app.listPurchasableBondsHandler))

	return app.enableCors(app.rateLimit(router))
}

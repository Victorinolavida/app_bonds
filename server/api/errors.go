package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {

	errMessage := map[string]any{
		"error": message,
	}

	err := app.WriteJson(w, status, errMessage)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}

}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "something went wrong. Please try again later."
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

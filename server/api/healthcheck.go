package main

import "net/http"

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status": "available",
	}

	err := app.WriteJson(w, http.StatusOK, data)
	if err != nil {

		app.serverErrorResponse(w, r, err)
		return
	}
}

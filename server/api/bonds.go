package main

import "net/http"

func (app *application) createBondHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		NumberBonds int    `json:"number_bonds"`
	}

	app.ReadJson(w, r, &input)

	app.WriteJson(w, http.StatusOK, envelop{"bond": input})
}
func (app *application) listAllBonds(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list all bounds"))
}

func (app *application) listAllBondsCreatedByUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list all bounds created by user"))
}
func (app *application) buyABondById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buy a bound by id"))
}

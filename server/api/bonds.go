package main

import (
	"net/http"

	"boundsApp.victorinolavida/internal/data"
	"boundsApp.victorinolavida/internal/validator"
)

func (app *application) createBondHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string     `json:"name"`
		Price       data.Price `json:"price"`
		NumberBonds int        `json:"number_bonds"`
	}

	app.ReadJson(w, r, &input)
	v := validator.New()

	bond := &data.Bond{
		Name:        input.Name,
		Price:       input.Price,
		NumberBonds: input.NumberBonds,
	}

	data.ValidateBond(v, bond)

	if !v.IsValid() {
		app.fieldValidationResponse(w, r, v.Errors)
		return
	}

	user := app.contextGetUser(r)

	newBond, err := app.models.Bonds.Insert(bond, user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.WriteJson(w, http.StatusOK, envelop{"bond": newBond})
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

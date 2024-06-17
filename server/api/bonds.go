package main

import (
	"errors"
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
func (app *application) listOwnedBondsLoggedUserHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list all bounds"))
}

func (app *application) listOwnedBondsLoggedUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	bonds, err := app.models.Bonds.GetBondsByUser(*user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.badRequestResponse(w, r, err)
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}
	app.WriteJson(w, http.StatusOK, envelop{"bonds": bonds})
}
func (app *application) buyABondByIdHandler(w http.ResponseWriter, r *http.Request) {
	bondId := app.readStringParamByName(r, "id")
	user := app.contextGetUser(r)
	if bondId == "" {
		app.badRequestResponse(w, r, errors.New("missing bond id"))
		return
	}

	bond := &data.Bond{ID: bondId}
	transaction := &data.Transaction{
		BondId:  bond.ID,
		BuyerID: user.ID,
	}

	err := app.models.Bonds.GetBondByID(bond)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.badRequestResponse(w, r, err)
			return
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	err = app.models.Bonds.IsPurchasableBound(bond, user, transaction)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrBoughtAlreadyBought):
			app.bondAlreadyOwnedResponse(w, r)
			return
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	newTransaction, error := app.models.Transactions.Insert(transaction)
	if error != nil {
		app.serverErrorResponse(w, r, error)
		return
	}

	app.WriteJson(w, http.StatusOK, envelop{"transaction": newTransaction})
}

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

	err := app.ReadJson(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	v := validator.New()

	user := app.contextGetUser(r)
	bond := &data.Bond{
		Name:        input.Name,
		Price:       input.Price,
		NumberBonds: input.NumberBonds,
		OwnerId:     user.ID,
		CreatedBy:   user.ID,
	}

	data.ValidateBond(v, bond)

	if !v.IsValid() {
		app.fieldValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Bonds.Insert(bond)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.WriteJson(w, http.StatusOK, envelop{"bond": bond})
}
func (app *application) listPurchasableBondsHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	var pagination data.Pagination
	query := r.URL.Query()

	v := validator.New()
	pagination.CurrentPage = app.readIntParamByName(query, "page", 1, v)
	pagination.PageSize = app.readIntParamByName(query, "page_size", 20, v)

	if data.ValidatePagination(v, &pagination); !v.IsValid() {
		app.fieldValidationResponse(w, r, v.Errors)
		return
	}

	bonds, paginationData, err := app.models.Bonds.GetAllAvailable(user, pagination)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.WriteJson(w, http.StatusOK, envelop{"bonds": bonds, "pagination": paginationData})
}

func (app *application) listOwnedBondsLoggedUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	query := r.URL.Query()
	var input struct {
		data.Pagination
	}
	v := validator.New()

	input.CurrentPage = app.readIntParamByName(query, "page", 1, v)
	input.PageSize = app.readIntParamByName(query, "page_size", 20, v)

	if data.ValidatePagination(v, &input.Pagination); !v.IsValid() {
		app.fieldValidationResponse(w, r, v.Errors)
		return
	}

	bonds, pagination, err := app.models.Bonds.GetBondsByUser(*user, input.Pagination)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.badRequestResponse(w, r, err)
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}
	app.WriteJson(w, http.StatusOK, envelop{"bonds": bonds, "pagination": pagination})
}
func (app *application) buyBondHandler(w http.ResponseWriter, r *http.Request) {
	bondId := app.readStringParamByName(r, "id")
	user := app.contextGetUser(r)
	if bondId == "" {
		app.badRequestResponse(w, r, errors.New("missing bond id"))
		return
	}

	bond := &data.Bond{ID: bondId}

	err := app.models.Bonds.GetBondByID(bond)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.badRequestResponse(w, r, errors.New("bond not found"))
			return
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}
	if bond.OwnerId == user.ID {
		app.bondAlreadyOwnedResponse(w, r)
		return
	}

	// creating transaction
	newTransaction := &data.Transaction{
		SellerID: bond.OwnerId,
		BuyerID:  user.ID,
		Price:    bond.Price,
	}

	//inserting transaction

	err = app.models.Transactions.Insert(newTransaction)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	bondTransaction := &data.BondTransaction{
		BondID:        bond.ID,
		TransactionID: newTransaction.ID,
	}
	err = app.models.BondTransaction.Insert(bondTransaction)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.models.Bonds.ChangeOwner(bond, user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.WriteJson(w, http.StatusOK, envelop{"transaction": newTransaction})
}

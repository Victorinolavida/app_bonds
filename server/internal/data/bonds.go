package data

import "boundsApp.victorinolavida/internal/validator"

// "boundsApp.victorinolavida/internal/validator"

type Bond struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	NumberBonds int    `json:"number_bonds"`
}

func ValidateBond(v *validator.Validator, bond *Bond) {
	v.Check(bond.Name != "", "name", "must be provided")
	v.Check(len(bond.Name) <= 40, "name", "must not be more than 40 characters long")
	v.Check(len(bond.Name) >= 3, "name", "must be at least 3 characters long")
	v.Check(bond.Price > 0, "price", "must be greater than zero")
	v.Check(bond.Price <= 100_000_000, "price", "must be less than or equal to 100000000")
	v.Check(bond.NumberBonds > 0, "number_bonds", "must be greater than zero")
	v.Check(bond.NumberBonds <= 10_000, "number_bonds", "must be less than or equal to 10000")
}

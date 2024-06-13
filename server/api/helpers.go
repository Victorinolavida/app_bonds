package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) WriteJson(w http.ResponseWriter, status int, data any) error {
	dataParsed, err := json.Marshal(data)

	if err != nil {

		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataParsed)

	return nil
}

func (app *application) ReadJson(w http.ResponseWriter, r *http.Request, destination any) error {

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(destination)
	if err != nil {

		return err
	}

	return nil
}

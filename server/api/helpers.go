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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"boundsApp.victorinolavida/internal/validator"
	"github.com/julienschmidt/httprouter"
)

type envelop map[string]any

var TokenName = "__auth_token"

var (
	ErrInvalidValue = errors.New("invalid value")
	ErrorNoParam    = errors.New("no param found")
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
		var syntaxError *json.SyntaxError
		var marshalError *json.UnmarshalTypeError

		switch {
		case errors.Is(err, syntaxError):
			return errors.New("syntax error")
		case errors.As(err, &marshalError):
			if marshalError.Field != "" {
				return fmt.Errorf("incorrect type for the field %s", marshalError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", marshalError.Offset)
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			field := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains a unknown key %s", field)
		case errors.Is(err, io.EOF):
			return fmt.Errorf("body can not be empty")

		default:
			return err
		}

	}

	return nil
}

func (app *application) getCookieByName(r *http.Request, name string) (string, error) {

	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}

	// Return the decoded cookie value.
	return string(cookie.Value), nil
}

func (app *application) readStringParamByName(r *http.Request, name string) string {
	params := httprouter.ParamsFromContext(r.Context())
	param := params.ByName(name)

	return param
}

func (app *application) readIntParamByName(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	paramString := qs.Get(key)
	if paramString == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(paramString)
	if err != nil {
		v.AddError(key, "must be a integer")
		return defaultValue
	}
	return i

}

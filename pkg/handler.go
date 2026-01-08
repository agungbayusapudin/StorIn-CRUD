package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
)

func DecodeResposne(r *http.Request, v interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("content-type harus berbentuk aplication/json")
	}

	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()
	if err := decode.Decode(v); err != nil {
		return err
	}
	return nil
}

func WriteResponse(w http.ResponseWriter, StatusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	return nil
}

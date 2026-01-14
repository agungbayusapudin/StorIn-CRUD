package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

func DecodeRequest(r *http.Request, v interface{}) error {
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

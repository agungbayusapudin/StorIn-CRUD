package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

func DecodeRequest(r *http.Request, data interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("content-type harus berbentuk aplication/json")
	}

	return json.NewDecoder(r.Body).Decode(data)
}

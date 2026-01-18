package response

import (
	"encoding/json"
	"net/http"
	"videocall/middleware"
)

// fotmat response
type StandardResponse struct {
	Response  interface{} `json:"response"`
	Metadata  Metadata    `json:"metadata"`
	ProccesId string      `json:"process_id"`
}

type Metadata struct {
	Message     string `json:"message"`
	Code        int    `json:"code"`
	ReponseCode string `json:"response_code"`
}

func WriteResponse(w http.ResponseWriter, r *http.Request, statusCode int, responseCode string, message string, data interface{}) {
	response := StandardResponse{
		Response: data,
		Metadata: Metadata{
			Message:     message,
			Code:        statusCode,
			ReponseCode: responseCode,
		},
		ProccesId: middleware.GetId(r.Context()),
	}

	value, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(value)
}

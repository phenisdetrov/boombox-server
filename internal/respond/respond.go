package respond

import (
	"encoding/json"
	"net/http"
)

func WriteWithStatus(w http.ResponseWriter, returnCode int, statusType, statusCode, statusMessage string, payload any) {
	type statusBody struct {
		Type    string `json:"type"`
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	type responseBody struct {
		Status  statusBody `json:"status"`
		Payload any        `json:"payload,omitempty"`
	}
	var response responseBody
	response.Status.Type, response.Status.Code, response.Status.Message, response.Payload = statusType, statusCode, statusMessage, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(returnCode)
	json.NewEncoder(w).Encode(response)
}

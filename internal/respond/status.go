package respond

import (
	"encoding/json"
	"net/http"
)

func WriteStatus(w http.ResponseWriter, httpCode int, typ, code, message string) {
	type status struct {
		Type    string `json:"type"`
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	type statusBody struct {
		Status status `json:"status"`
	}
	var response statusBody
	response.Status.Type, response.Status.Code, response.Status.Message = typ, code, message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}

package check

import (
	"encoding/json"
	"net/http"

	"github.com/phenisdetrov/boombox-server/internal/config"
)

func StatusGet(w http.ResponseWriter, r *http.Request) {
	type statusGetResponse struct {
		Status  string `json:"status"`
		Version string `json:"version"`
	}
	var response statusGetResponse
	response.Status = "OK"
	response.Version = config.VERSION
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

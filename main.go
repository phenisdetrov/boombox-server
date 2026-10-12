package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const listenAddr = "10.0.0.1:20000"
const version = "0.0.1"

func main() {
	http.HandleFunc("GET /api/v1/status", statusGet)

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func statusGet(w http.ResponseWriter, r *http.Request) {
	type statusGetResponse struct {
		Status  string `json:"status"`
		Version string `json:"version"`
	}
	var response statusGetResponse
	response.Status = "OK"
	response.Version = version
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

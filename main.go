package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const listenAddr = "10.0.0.1:20000"
const version = "0.0.1"

var adminUsername string
var adminPassword string

func main() {
	http.HandleFunc("GET /api/v1/status", statusGet)
	http.HandleFunc("POST /api/v1/admin/create", adminCreatePost)

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

func adminCreatePost(w http.ResponseWriter, r *http.Request) {
	type adminCreatePostRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type adminCreatePostResponse struct {
		Type    string `json:"type"`
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	var request adminCreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	var response adminCreatePostResponse
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Type, response.Code, response.Message = "KO", "BAD_REQUEST", "Request expressed not according to contract"
		json.NewEncoder(w).Encode(response)
		return
	}
	if adminUsername != "" {
		w.WriteHeader(http.StatusConflict)
		response.Type, response.Code, response.Message = "KO", "ADMIN_EXISTS", "Admin account already created"
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response.Type, response.Code, response.Message = "OK", "SUCCESS", "Admin account successfully created"
	json.NewEncoder(w).Encode(response)
	adminUsername = request.Username
	adminPassword = request.Password
}

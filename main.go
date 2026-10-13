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

type status struct {
	Type    string `json:"type"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
type statusBody struct {
	Status status `json:"status"`
}

func main() {
	http.HandleFunc("GET /api/v1/status", statusGet)
	http.HandleFunc("POST /api/v1/admin/create", adminCreatePost)

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func writeStatus(w http.ResponseWriter, httpCode int, typ, code, message string) {
	var response statusBody
	response.Status.Type, response.Status.Code, response.Status.Message = typ, code, message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
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
	var request adminCreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeStatus(w, http.StatusBadRequest, "KO", "BAD_REQUEST", "Request expressed not according to contract")
		return
	}
	if adminUsername != "" {
		writeStatus(w, http.StatusConflict, "KO", "ADMIN_EXISTS", "Admin account already created")
		return
	}
	adminUsername, adminPassword = request.Username, request.Password
	writeStatus(w, http.StatusCreated, "OK", "SUCCESS", "Admin account successfully created")
}

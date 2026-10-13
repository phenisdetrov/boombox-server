package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const listenAddr = "10.0.0.1:20000"
const version = "0.0.1"

var usernameRegexp = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)
var passwordRegexp = regexp.MustCompile(`^[\x20-\x7E]+$`)

var adminUsername string
var adminPasswordHash []byte
var adminCreated bool

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

func helperWriteStatus(w http.ResponseWriter, httpCode int, typ, code, message string) {
	var response statusBody
	response.Status.Type, response.Status.Code, response.Status.Message = typ, code, message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}

func helperValidCredentials(username, password string) bool {
	if !usernameRegexp.MatchString(username) || !passwordRegexp.MatchString(password) {
		return false
	}
	if len(username) > 72 || len(password) > 72 {
		return false
	}
	return true
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
	jsonParseErr := json.NewDecoder(r.Body).Decode(&request)

	if jsonParseErr != nil {
		helperWriteStatus(w, http.StatusBadRequest, "KO", "BAD_REQUEST", "Request expressed not according to contract")
		return
	}
	if !helperValidCredentials(request.Username, request.Password) {
		helperWriteStatus(w, http.StatusBadRequest, "KO", "BAD_REQUEST", "Request expressed not according to contract")
		return
	}
	if adminCreated {
		helperWriteStatus(w, http.StatusConflict, "KO", "ADMIN_EXISTS", "Admin account already created")
		return
	}
	hash, bcryptHashGenErr := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if bcryptHashGenErr != nil {
		helperWriteStatus(w, http.StatusInternalServerError, "KO", "INTERNAL_ERR", "Internal server error during managing admin credentials")
		return
	}
	adminCreated, adminUsername, adminPasswordHash = true, request.Username, hash
	helperWriteStatus(w, http.StatusCreated, "OK", "SUCCESS", "Admin account successfully created")
}

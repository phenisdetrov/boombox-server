package admin

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/phenisdetrov/boombox-server/internal/respond"
)

var adminUsername string
var adminPasswordHash []byte
var adminCreated bool

func Create(w http.ResponseWriter, r *http.Request) {
	type adminCreatePostRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var request adminCreatePostRequest
	jsonParseErr := json.NewDecoder(r.Body).Decode(&request)

	if jsonParseErr != nil {
		respond.WriteStatus(w, http.StatusBadRequest, "KO", "BAD_REQUEST", "Request expressed not according to contract")
		return
	}
	if !checkCredentials(request.Username, request.Password) {
		respond.WriteStatus(w, http.StatusBadRequest, "KO", "BAD_REQUEST", "Request expressed not according to contract")
		return
	}
	if adminCreated {
		respond.WriteStatus(w, http.StatusConflict, "KO", "ADMIN_EXISTS", "Admin account already created")
		return
	}
	hash, bcryptHashGenErr := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if bcryptHashGenErr != nil {
		respond.WriteStatus(w, http.StatusInternalServerError, "KO", "INTERNAL_ERR", "Internal server error during managing admin credentials")
		return
	}
	adminUsername, adminPasswordHash, adminCreated = request.Username, hash, true
	respond.WriteStatus(w, http.StatusCreated, "OK", "SUCCESS", "Admin account successfully created")
}

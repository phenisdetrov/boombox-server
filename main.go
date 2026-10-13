package main

import (
	"log"
	"net/http"

	"github.com/phenisdetrov/boombox-server/internal/get"

	"github.com/phenisdetrov/boombox-server/internal/check"
	"github.com/phenisdetrov/boombox-server/internal/admin"
)

var adminUsername string
var adminPasswordHash []byte
var adminCreated bool

func main() {
	http.HandleFunc("GET /api/v1/status", check.StatusGet)
	http.HandleFunc("POST /api/v1/admin/create", admin.Create)

	log.Fatal(http.ListenAndServe(get.LISTEN, nil))
}

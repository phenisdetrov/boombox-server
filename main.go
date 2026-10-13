package main

import (
	"log"
	"net/http"

	"github.com/phenisdetrov/boombox-server/internal/config"
	"github.com/phenisdetrov/boombox-server/internal/admin"
	"github.com/phenisdetrov/boombox-server/internal/check"
)

func main() {
	http.HandleFunc("GET /api/v1/status", check.StatusGet)
	http.HandleFunc("POST /api/v1/admin/create", admin.Create)

	log.Fatal(http.ListenAndServe(config.LISTEN, nil))
}

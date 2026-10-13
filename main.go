package main

import (
	"log"
	"net/http"

	"github.com/phenisdetrov/boombox-server/internal/config"
	"github.com/phenisdetrov/boombox-server/internal/endpoints/admin"
	"github.com/phenisdetrov/boombox-server/internal/endpoints/root"
)

func main() {
	http.HandleFunc("GET /api/v1/status", root.Status)
	http.HandleFunc("POST /api/v1/admin/create", admin.Create)

	log.Fatal(http.ListenAndServe(config.LISTEN, nil))
}

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phenisdetrov/boombox-server/internal/config"
	"github.com/phenisdetrov/boombox-server/internal/endpoints/admin"
	"github.com/phenisdetrov/boombox-server/internal/endpoints/root"
)

func main() {
	pool, err := pgxpool.New(context.Background(), config.POSTGRES)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /api/v1/status", root.Status(pool))
	http.HandleFunc("POST /api/v1/admin/create", admin.Create)

	log.Fatal(http.ListenAndServe(config.LISTEN, nil))
}

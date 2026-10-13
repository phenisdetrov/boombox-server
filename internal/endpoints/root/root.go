package root

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phenisdetrov/boombox-server/internal/config"
	"github.com/phenisdetrov/boombox-server/internal/respond"
)

func Status(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := struct {
			Version string `json:"version"`
		}{config.VERSION}
		err := pool.Ping(r.Context())
		if err != nil {
			respond.WriteWithStatus(w, http.StatusServiceUnavailable, "KO", "DB_UNAVAILABLE", "Server is not ready, Postgres unavailable", payload)
			return
		}
		respond.WriteWithStatus(w, http.StatusOK, "OK", "SUCCESS", "Server is ready, postgres available", payload)
	}
}

package root

import (
	"net/http"

	"github.com/phenisdetrov/boombox-server/internal/config"
	"github.com/phenisdetrov/boombox-server/internal/respond"
)

func Status(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Version string `json:"version"`
	}{config.VERSION}
	respond.WriteWithStatus(w, http.StatusOK, "OK", "SUCCESS", "Server is ready", payload)
}

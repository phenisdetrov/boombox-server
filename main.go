package main

import (
    "log"
    "net/http"
    "fmt"
)

const listenAddr = "10.0.0.1:20000"

func main () {
    http.HandleFunc("GET /api/v1/check", healthCheck)

    log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func healthCheck (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "ok\n")
}

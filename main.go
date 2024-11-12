package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

const port = "8080"

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}
	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", healthcheck)
	mux.HandleFunc("GET /admin/metrics", apiCfg.showrequests)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetrequests)
	mux.HandleFunc("POST /api/validate_chirp", validateChirp)
	// mux.HandleFunc("POST /api/validate_chirp")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}

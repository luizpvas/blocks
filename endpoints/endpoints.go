package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/luizpvas/blocks/config"
)

// StartHTTPServer starts the HTTP server.
func StartHTTPServer(app *config.AppConfig) error {
	log.Println("Starting HTTP server...")

	// Static files bundled with the binary distribution
	box := packr.NewBox("../client")
	http.Handle("/blocks/client/", http.StripPrefix("/blocks/client/", http.FileServer(box)))

	// Custom static files directory used for development.
	if app.HTTP.StaticFilesDir != "" {
		log.Printf("Serving static files from: %v\n", app.HTTP.StaticFilesDir)
		http.Handle("/", http.FileServer(http.Dir(app.HTTP.StaticFilesDir)))
	}

	http.HandleFunc("/app_config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	})

	http.HandleFunc("/resource_config", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		resource, ok := app.Resources[id]
		if !ok {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resource)
	})

	log.Printf("HTTP server started at [%s]", app.HTTP.Listen)

	return http.ListenAndServe(app.HTTP.Listen, nil)
}

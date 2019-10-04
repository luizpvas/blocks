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

	box := packr.NewBox("../client")
	http.Handle("/blocks/client/", http.StripPrefix("/blocks/client/", http.FileServer(box)))

	http.HandleFunc("/app_config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	})

	log.Printf("HTTP server started at [%s]", app.HTTP.Listen)

	return http.ListenAndServe(app.HTTP.Listen, nil)
}

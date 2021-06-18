//go:generate goversioninfo -icon=Logo.ico

package main

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("%s", r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	log.Info().Msg("reflector-server is starting...")
	hub := newRouter()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to respond to WebSocket requests")
	}
}

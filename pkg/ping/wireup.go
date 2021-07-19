package ping

import (
	"fmt"
	"log"
	"net/http"

	"goapi/internal/server"
)

func getPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping")
}

func New(s *server.Server) {
	log.Println("Ping wireup")
	s.Router.HandleFunc("/", getPing).Methods("GET")
	s.Router.HandleFunc("/ping", getPing).Methods("GET")
}

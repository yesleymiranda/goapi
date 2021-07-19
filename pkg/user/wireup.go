package user

import (
	"fmt"
	"log"
	"net/http"

	"goapi/internal/server"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user")
}

func New(s *server.Server) {
	log.Println("User wireup")
	s.Router.HandleFunc("/user", getUser).Methods("GET")
}

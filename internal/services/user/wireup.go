package user

import (
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/webapplication"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user")
}

func New(s *webapplication.App) {
	logger.Info("User wireup")
	s.Router.HandleFunc("/user", getUser).Methods("GET")
}

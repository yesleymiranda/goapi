package ping

import (
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/webapplication"
	"net/http"
)

func getPing(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "ping")
}

func New(s *webapplication.App) {
	logger.Info("Ping wireup")
	s.Router.HandleFunc("/", getPing).Methods("GET")
	s.Router.HandleFunc("/ping", getPing).Methods("GET")
}

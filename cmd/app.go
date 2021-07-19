package app

import (
	"net/http"

	"goapi/internal/server"
	"goapi/pkg/ping"
	"goapi/pkg/user"

	"github.com/gorilla/mux"
	logrus "github.com/sirupsen/logrus"
)

func Start() {
	logrus.Info("Starting service...")

	r := mux.NewRouter().StrictSlash(true)

	s := &server.Server{
		Router: r,
		Logger: *logrus.Logger,
	}

	ping.New(s)
	user.New(s)

	logrus.Fatal(http.ListenAndServe(":8000", r))
}

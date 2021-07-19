package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Router *mux.Router
	Logger logrus.Logger
}

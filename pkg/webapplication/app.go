package webapplication

import (
	"goapi/pkg/logger"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() (*App, error) {
	logger.Init()

	logger.Info("Start app...")

	r := mux.NewRouter().StrictSlash(true)

	newApp := &App{
		Router: r,
	}

	return newApp, nil
}

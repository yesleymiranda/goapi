package user

import (
	"github.com/yesleymiranda/go-toolkit/webapplication"
)

func New(app *webapplication.App) {
	repository := NewRepository("db-name")
	service := NewService(repository)

	routes(app, service)
}

func routes(app *webapplication.App, svc *Service) {
	app.Router.HandleFunc("/users/{id:[0-9]{1,12}}", MakeGetByIDHandler(svc)).Methods("GET")
}

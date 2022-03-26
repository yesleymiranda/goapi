package app

import (
	"goapi/app/user"

	"github.com/yesleymiranda/go-toolkit/webapplication"
)

const port = "8080"

func Run() error {
	app := webapplication.New(&webapplication.ApplicationConfig{
		Port:     port,
		WithPing: true,
	})
	app.Initialize()

	user.New(app)

	return app.ListenAndServe()
}

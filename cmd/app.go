package app

import (
	"goapi/api/serve/ping"
	"goapi/api/serve/user"
	"goapi/pkg/webapplication"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Start() {
	app, err := webapplication.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Error on start web application")
	}

	ping.New(app)
	user.New(app)

	log.Fatal().Err(http.ListenAndServe(":8000", app.Router)).Msg("Error on listen and serve")
}

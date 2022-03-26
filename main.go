package main

import (
	"errors"
	"goapi/app"
	"net/http"
)

func main() {
	err := app.Run()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

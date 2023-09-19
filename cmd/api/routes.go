package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/v1/healthcheck", app.healthcheckHandler)

	return r
}

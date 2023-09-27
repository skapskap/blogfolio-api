package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.NotFound(app.notFoundResponse)

	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Get("/v1/healthcheck", app.healthcheckHandler)
	r.Post("/v1/posts", app.createPostHandler)
	r.Get("/v1/posts/{id}", app.showPostHandler)
	r.Put("/v1/posts/{id}", app.updatePostHandler)
	r.Delete("/v1/posts/{id}", app.deletePostHandler)

	return r
}

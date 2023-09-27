package main

import (
	"fmt"
	"github.com/skapskap/blogfolio-api/internal/data"
	"net/http"
	"time"
)

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	post := &data.Post{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	err = app.models.Posts.Insert(post)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posts/%d", post.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"post": post}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showPostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	post := data.Post{
		ID:          id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		PublishedAt: time.Now(),
		Title:       "Construí este portfolio com SvelteKit + Go + TailwindCSS em 3 dias!",
		Description: "Aqui fica o conteúdo do blog",
		Status:      "Publicado",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"post": post}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "O servidor encontrou um problema e não processou sua requisição", http.StatusInternalServerError)
	}
}

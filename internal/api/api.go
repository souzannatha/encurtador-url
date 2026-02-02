package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

type apiResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/url", func(r chi.Router) {
			r.Post("/shorten", handleShortenURL(db))
			r.Get("/{code}", handleGetShortenURL(db))
		})
	})
	return r
}

type PostBody struct {
	URL string `json:"url"`
}

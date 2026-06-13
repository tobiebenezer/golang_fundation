package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PostHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

func NewPostRouter(h PostHandler, paginateMiddleware func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.Create)
	r.With(paginateMiddleware).Get("/", h.List)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Patch("/", h.Update)
		r.Delete("/", h.Delete)
	})

	return r
}

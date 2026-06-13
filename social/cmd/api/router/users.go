package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

func NewUserRouter(h UserHandler, paginateMiddleware func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.Register)
	r.With(paginateMiddleware).Get("/", h.List)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Patch("/", h.Update)
		r.Delete("/", h.Delete)
	})

	return r
}

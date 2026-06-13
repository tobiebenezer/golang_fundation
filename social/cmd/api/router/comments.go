package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CommentHandler interface {
	CreateComment(w http.ResponseWriter, r *http.Request)
	GetCommentByID(w http.ResponseWriter, r *http.Request)
	UpdateComment(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)
	ListCommentsByPostID(w http.ResponseWriter, r *http.Request)
	ListCommentsByUserID(w http.ResponseWriter, r *http.Request)
	CountCommentsByPostID(w http.ResponseWriter, r *http.Request)
	CountCommentsByUserID(w http.ResponseWriter, r *http.Request)
}

func NewCommentRouter(h CommentHandler, paginateMiddleware func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.CreateComment)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetCommentByID)
		r.Patch("/", h.UpdateComment)
		r.Delete("/", h.DeleteComment)
	})

	r.Route("/posts/{postID}", func(r chi.Router) {
		r.With(paginateMiddleware).Get("/", h.ListCommentsByPostID)
		r.Get("/count", h.CountCommentsByPostID)
	})

	r.Route("/users/{userID}", func(r chi.Router) {
		r.With(paginateMiddleware).Get("/", h.ListCommentsByUserID)
		r.Get("/count", h.CountCommentsByUserID)
	})

	return r
}

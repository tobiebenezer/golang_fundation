package main

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	customMiddleware "social/update/internal/middleware"
	"social/update/internal/domain"
	"social/update/cmd/api/router"
	"social/update/docs"
)

type application struct {
	config      config
	userService domain.UserService
	postService domain.PostService
	commentService domain.CommentService
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	addr string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime string
	
}

func (a *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json", "text/xml"))
	r.Use(middleware.NoCache)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() t hat the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	uri, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	r.Get("/swagger/*", docs.SwaggerHandler(uri))
	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})	

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.healthCheckHandler)

		r.Mount("/users", router.NewUserRouter(a, customMiddleware.Paginate))
		r.Mount("/posts", router.NewPostRouter(&postHandler{app: a}, customMiddleware.Paginate))
		r.Mount("/comments", router.NewCommentRouter(NewCommentHandler(a), customMiddleware.Paginate))
	})

	return r
}

func (a *application) initServer(r http.Handler) error {

	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      r,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Println("starting server on", a.config.addr)

	return srv.ListenAndServe()
}

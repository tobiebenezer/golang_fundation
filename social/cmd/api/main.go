package main

import (
	"log"
	"social/update/internal/db"
	"social/update/internal/env"
	"social/update/internal/service"
	"social/update/internal/store"
)


//	@title			GopherSocial API
//	@version		1.0
//	@description	This is a social media application API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1
func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://postgres:Awodumila@localhost:5432/gosocial?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.NewDB(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic("failed to connect to database: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	log.Println("connected to database successfully")

	userRepo := store.NewUserRepository(db)
	postRepo := store.NewPostRepository(db)
	commentRepo := store.NewCommentRepository(db)

	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)
	commentService := service.NewCommentService(commentRepo, postRepo, userRepo)


	app := &application{
		config:      cfg,
		userService: userService,
		postService: postService,
		commentService: commentService,
	}


	log.Fatal(app.initServer(app.mount()))

}

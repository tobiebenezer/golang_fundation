package main

import (
	"net/http"
)

// type app struct {
// 	addr string
// }


// func (s *app) getUsersHandler(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("all users"))
// }

// func (s *app) createUsersHandler(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("user created"))
// }

func main() {
	app := &api{addr: ":8081"}

	//Initialise the serveMux()
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: app.addr,
		Handler: mux,
	}

	mux.HandleFunc("Get /user", app.getUsersHandler)
	mux.HandleFunc("POST /user", app.createUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

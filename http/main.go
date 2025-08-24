package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All users"))
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created user"))
}

func main() {
	api := &api{addr: ":8080"}

	// Initialize the ServeMux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	log.Printf("Server running on %s", api.addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

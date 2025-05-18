package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main(){
	r := chi.NewRouter();
	r.Use(middleware.Logger);
	r.Use(middleware.Recoverer);
	r.Use(cors.AllowAll().Handler);//tempoarily allow all origins

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Go server is active!!!"))
	})

	http.ListenAndServe(":8000",r)
}
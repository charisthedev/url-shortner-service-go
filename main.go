package main

import (
	"log"
	"net/http"
	"os"
	dbconfig "url-shortner/db/dbConfig"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	portString := os.Getenv("PORT");

	db := dbconfig.ConnectDb();
	defer db.Close()


	r := chi.NewRouter();
	r.Use(middleware.Logger);
	r.Use(middleware.Recoverer);
	r.Use(cors.AllowAll().Handler);//tempoarily allow all origins

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Go server is active!!!"))
	})

	srv := http.Server{
		Addr: ":" + portString,
		Handler: r,
	}

	func()  {
		log.Printf("server is running on port %s",portString);
		err = srv.ListenAndServe();
		if err != nil {
			log.Fatal(err)
		}	
	}()
}
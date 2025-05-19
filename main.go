package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main(){
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	dbStr := os.Getenv("DB_URL");
	portString := os.Getenv("PORT");

	db, err := sql.Open("postgres",dbStr);

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close();

	if err = db.Ping(); err != nil {
		log.Fatal((err))
	}

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
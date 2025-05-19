package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main(){
	err := godotenv.Load("../.env");
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	dbStr := os.Getenv("DB_URL");

	db, err := sql.Open("postgres",dbStr);

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal((err))
	}

	defer db.Close();

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create DB driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Migration setup error: %v", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migrations to run.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		log.Println("Migrations ran successfully.")
	}
}
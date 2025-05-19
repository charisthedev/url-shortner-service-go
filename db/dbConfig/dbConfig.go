package dbconfig

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
);

func ConnectDb() *sql.DB{
	dbStr := os.Getenv("DB_URL");

	db, err := sql.Open("postgres",dbStr);

	if err != nil {
		log.Fatal("Failed to connect to db",err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal((err))
	}

	fmt.Println("Connected to DB");

	return db
}
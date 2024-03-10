package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"library-app/pkg/library-app/model"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
}

func main() {
	fmt.Println("Started server")
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8080", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:12345@localhost/library_app?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: model.NewModels(db),
	}

	app.run()
}

func (app *application) run() {
	fmt.Println("Running")
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Users
	v1.HandleFunc("/register", app.createUserHandler).Methods("POST")
	v1.HandleFunc("/users", app.getAllUsersHandler).Methods("GET")
	v1.HandleFunc("/users/{userId:[0-9]+}", app.getUserByIdHandler).Methods("GET")
	v1.HandleFunc("/users/{userId:[0-9]+}", app.updateUserHandler).Methods("PUT")
	v1.HandleFunc("/users/{userId:[0-9]+}", app.deleteUserHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

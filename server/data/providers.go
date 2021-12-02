package data

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/wire"
)

func ProvideRepository(db *sql.DB) *Repository {
	return &Repository{ Db: db }
}

func ProvideHttpHandler(repository *Repository) http.HandlerFunc {
	return HttpHandler(repository)
}

func ProvideDbConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/forum?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		db.Close()
	}
	return db, err
	// return nil, errors.New("DB: Not Implemented")
}

var Providers = wire.NewSet(ProvideRepository, ProvideHttpHandler, ProvideDbConnection)

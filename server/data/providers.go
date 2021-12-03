package data

import (
	"SamuraiLab3/server/db"
	"database/sql"
	"net/http"

	"github.com/google/wire"
)

func ProvideRepository(db *sql.DB) *Repository {
	return &Repository{Db: db}
}

func ProvideHttpHandler(repository *Repository) http.HandlerFunc {
	return HttpHandler(repository)
}

func ProvideDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "forum",
		User:       "postgres",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

var Providers = wire.NewSet(ProvideRepository, ProvideHttpHandler, ProvideDbConnection)

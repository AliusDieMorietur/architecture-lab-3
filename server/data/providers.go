package data

import (
	"database/sql"
	"errors"
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
	// conn := &db.Connection{
	// 	DbName:     "chat-example",
	// 	User:       "roman",
	// 	Host:       "localhost",
	// 	DisableSSL: true,
	// }
	// return conn.Open()
	return nil, errors.New("DB: Not Implemented")
}

var Providers = wire.NewSet(ProvideRepository, ProvideHttpHandler, ProvideDbConnection)

package data

import (
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

var Providers = wire.NewSet(ProvideRepository, ProvideHttpHandler)

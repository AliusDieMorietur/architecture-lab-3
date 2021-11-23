package data

import (
	"database/sql"

	"github.com/google/wire"
)

func ProvideRepository(db *sql.DB) *Repository {
	return &Repository{Db: db}
}

// TODO Add http provider

var Providers = wire.NewSet(ProvideRepository)

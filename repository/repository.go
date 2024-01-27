// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}
	return &Repository{
		Db: db,
	}
}

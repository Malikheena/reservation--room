package dbrepo

import (
	"database/sql"

	"github.com/malikheena/heena2-main/internal/config"
	"github.com/malikheena/heena2-main/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

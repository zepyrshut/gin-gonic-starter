package dbrepo

import (
	"database/sql"

	"github.com/zepyrshut/gin-gonic-starter/internal/config"
	"github.com/zepyrshut/gin-gonic-starter/internal/repository"
)

type postgreDBRepo struct {
	App *config.Application
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, app *config.Application) repository.DBRepo {
	return &postgreDBRepo{
		App: app,
		DB:  conn,
	}
}

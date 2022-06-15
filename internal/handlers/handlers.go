package handlers

import (
	"github.com/zepyrshut/gin-gonic-starter/internal/config"
	"github.com/zepyrshut/gin-gonic-starter/internal/database"
	"github.com/zepyrshut/gin-gonic-starter/internal/repository"
	"github.com/zepyrshut/gin-gonic-starter/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.Application
	DB  repository.DBRepo
}

func NewRepo(a *config.Application, db *database.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

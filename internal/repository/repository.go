package repository

import "github.com/zepyrshut/gin-gonic-starter/internal/models"

type DBRepo interface {
	OneUser(id int) (models.User, error)
}

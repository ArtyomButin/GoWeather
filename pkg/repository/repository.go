package repository

import (
	"github.com/ArtyomButin/GoWeather/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *pgxpool.Pool) *Repository  {
	return &Repository{
		Authorization: NewAuthUser(db),
	}
}
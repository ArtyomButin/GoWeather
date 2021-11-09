package services

import (
	"github.com/ArtyomButin/GoWeather/internal/models"
	"github.com/ArtyomButin/GoWeather/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
	}

}

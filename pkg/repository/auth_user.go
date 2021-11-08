package repository

import (
	"context"
	"fmt"
	"github.com/ArtyomButin/GoWeather/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthUser struct {
	db *pgxpool.Pool
}

func NewAuthUser(db *pgxpool.Pool) *AuthUser {
	return &AuthUser{db: db}
}

func (r *AuthUser) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, email, password_hash) values ($1, $2, $3, $4) RETURNING id",
		usersTable,
	)
	row := r.db.QueryRow(context.Background(), query, user.Name, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthUser) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	var users []*models.User
	err := pgxscan.Select(context.Background(), r.db, &users, query, username, password)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}

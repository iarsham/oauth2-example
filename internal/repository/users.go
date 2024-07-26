package repository

import (
	"context"
	"database/sql"
	"github.com/iarsham/oauth2-example/internal/domain"
	"github.com/iarsham/oauth2-example/internal/entities"
	"github.com/iarsham/oauth2-example/internal/models"
	"time"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UsersRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (a *userRepositoryImpl) FindByEmail(email string) (*models.Users, error) {
	query := `SELECT * FROM users WHERE email = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row := a.db.QueryRowContext(ctx, query, email)
	var user models.Users
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Picture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *userRepositoryImpl) Create(data *entities.GoogleOAuthResponse) (*models.Users, error) {
	query := `INSERT INTO users (name, email, picture) VALUES ($1, $2, $3) RETURNING *`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	args := []any{data.Name, data.Email, data.Picture}
	row := a.db.QueryRowContext(ctx, query, args...)
	var user models.Users
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Picture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

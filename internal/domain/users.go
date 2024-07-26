package domain

import (
	"github.com/iarsham/oauth2-example/internal/entities"
	"github.com/iarsham/oauth2-example/internal/models"
)

type UsersRepository interface {
	FindByEmail(email string) (*models.Users, error)
	Create(data *entities.GoogleOAuthResponse) (*models.Users, error)
}

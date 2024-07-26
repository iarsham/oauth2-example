package domain

import (
	"context"
	"github.com/iarsham/oauth2-example/internal/entities"
	"github.com/iarsham/oauth2-example/internal/models"
	"golang.org/x/oauth2"
)

type AuthService interface {
	RetrieveOAuth2Token(code string) (*oauth2.Token, error)
	RetrieveUserDataByToken(token *oauth2.Token) ([]byte, error)
	UnmarshalUserData(data []byte) (*entities.GoogleOAuthResponse, error)
	FindUserByEmail(email string) (*models.Users, error)
	CreateUser(data *entities.GoogleOAuthResponse) (*models.Users, error)
	PutSession(ctx context.Context, value any)
	PopSession(ctx context.Context)
}
		
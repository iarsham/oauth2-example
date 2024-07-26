package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/alexedwards/scs/v2"
	"github.com/iarsham/oauth2-example/configs"
	"github.com/iarsham/oauth2-example/internal/domain"
	"github.com/iarsham/oauth2-example/internal/entities"
	"github.com/iarsham/oauth2-example/internal/models"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
)

type authServiceImpl struct {
	db        *sql.DB
	logger    *zap.Logger
	session   *scs.SessionManager
	userRepo  domain.UsersRepository
	oauth2Cfg *oauth2.Config
}

func NewAuthService(userRepo domain.UsersRepository, db *sql.DB, logger *zap.Logger, session *scs.SessionManager, cfg *configs.Config) domain.AuthService {
	return &authServiceImpl{
		db:       db,
		logger:   logger,
		userRepo: userRepo,
		session:  session,
		oauth2Cfg: &oauth2.Config{
			ClientID:     cfg.App.GoogleClientId,
			ClientSecret: cfg.App.GoogleClientSecret,
			RedirectURL:  cfg.App.GoogleRedirectUrl,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint: google.Endpoint,
		},
	}
}

func (a *authServiceImpl) RetrieveOAuth2Token(code string) (*oauth2.Token, error) {
	token, err := a.oauth2Cfg.Exchange(context.Background(), code)
	if err != nil {
		a.logger.Error("Error exchanging code", zap.Error(err))
		return nil, err
	}
	return token, nil
}

func (a *authServiceImpl) RetrieveUserDataByToken(token *oauth2.Token) ([]byte, error) {
	request, err := a.oauth2Cfg.Client(context.Background(), token).Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		a.logger.Error("Error getting user info from google token", zap.Error(err))
		return nil, err
	}
	defer request.Body.Close()
	data, err := io.ReadAll(request.Body)
	if err != nil {
		a.logger.Error("Error reading user info from google request body", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (a *authServiceImpl) UnmarshalUserData(data []byte) (*entities.GoogleOAuthResponse, error) {
	var resp entities.GoogleOAuthResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		a.logger.Error("Error unmarshalling user data", zap.Error(err))
		return nil, err
	}
	return &resp, nil
}

func (a *authServiceImpl) FindUserByEmail(email string) (*models.Users, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		a.logger.Error("Error finding user by email", zap.Error(err))
		return nil, err
	}
	return user, nil
}
func (a *authServiceImpl) CreateUser(data *entities.GoogleOAuthResponse) (*models.Users, error) {
	user, err := a.userRepo.Create(data)
	if err != nil {
		a.logger.Error("Error creating user", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (a *authServiceImpl) PutSession(ctx context.Context, value any) {
	a.session.Put(ctx, "session_id", value)
}

func (a *authServiceImpl) PopSession(ctx context.Context) {
	a.session.PopString(ctx, "session_id")
}

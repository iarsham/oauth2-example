package routers

import (
	"database/sql"
	"github.com/iarsham/multiplexer"
	"github.com/iarsham/oauth2-example/configs"
	"github.com/iarsham/oauth2-example/internal/handlers"
	"github.com/iarsham/oauth2-example/internal/middlewares"
	"github.com/iarsham/oauth2-example/internal/repository"
	"github.com/iarsham/oauth2-example/internal/services"
	"github.com/iarsham/oauth2-example/pkg/session"
	"go.uber.org/zap"
	"net/http"
)

func New(db *sql.DB, logger *zap.Logger, cfg *configs.Config) http.Handler {
	r := multiplexer.New(http.NewServeMux(), "/api")
	scs := session.New(cfg)
	userRepo := repository.NewUserRepository(db)
	oauthHandler := handlers.AuthHandler{
		Service: services.NewAuthService(userRepo, db, logger, scs, cfg),
	}
	protected := multiplexer.NewChain(middlewares.AuthMiddleware(scs))
	r.Handle("GET /user", protected.WrapFunc(oauthHandler.GetUserByEmailHandler))
	r.Handle("GET /logout", protected.WrapFunc(oauthHandler.LogoutUserHandler))
	r.HandleFunc("POST /login/google", oauthHandler.LoginOAuth2GoogleHandler)
	return middlewares.CorsMiddleware(cfg).Handler(scs.LoadAndSave(r))
}

package session

import (
	"github.com/alexedwards/scs/v2"
	"github.com/iarsham/oauth2-example/configs"
)

func New(cfg *configs.Config) *scs.SessionManager {
	session := scs.New()
	session.Cookie.Name = "session_id"
	session.Cookie.Secure = !cfg.App.Debug
	return session
}

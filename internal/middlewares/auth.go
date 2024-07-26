package middlewares

import (
	"context"
	"github.com/alexedwards/scs/v2"
	"github.com/iarsham/bindme"
	"net/http"
)

func AuthMiddleware(scsManager *scs.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			email := scsManager.GetString(r.Context(), "session_id")
			if email == "" {
				bindme.WriteJson(w, http.StatusUnauthorized, nil, nil)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), "email", email))
			next.ServeHTTP(w, r)
		})
	}
}

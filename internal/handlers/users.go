package handlers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/bindme"
	"net/http"
)

func (a *AuthHandler) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)
	user, err := a.Service.FindUserByEmail(email)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			bindme.WriteJson(w, http.StatusNotFound, nil, nil)
		case err != nil:
			bindme.WriteJson(w, http.StatusInternalServerError, nil, nil)
		}
		return
	}
	bindme.WriteJson(w, http.StatusOK, user, nil)
}

func (a *AuthHandler) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	a.Service.PopSession(r.Context())
	bindme.WriteJson(w, http.StatusOK, nil, nil)
}

package handlers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/bindme"
	"github.com/iarsham/oauth2-example/internal/domain"
	"github.com/iarsham/oauth2-example/internal/entities"
	"net/http"
)

type M map[string]any
type AuthHandler struct {
	Service domain.AuthService
}

func (a *AuthHandler) LoginOAuth2GoogleHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.GoogleOAuthRequest)
	if err := bindme.ReadJson(r, data); err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, M{"error": err.Error()}, nil)
		return
	}
	token, err := a.Service.RetrieveOAuth2Token(data.Code)
	if err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, M{"error": err.Error()}, nil)
		return
	}
	userData, err := a.Service.RetrieveUserDataByToken(token)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, M{"error": http.StatusText(http.StatusInternalServerError)}, nil)
		return
	}
	userObj, err := a.Service.UnmarshalUserData(userData)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, M{"error": http.StatusText(http.StatusInternalServerError)}, nil)
		return
	}
	user, err := a.Service.FindUserByEmail(userObj.Email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		user, err = a.Service.CreateUser(userObj)
		if err != nil {
			bindme.WriteJson(w, http.StatusInternalServerError, M{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
	}
	a.Service.PutSession(r.Context(), user.Email)
	bindme.WriteJson(w, http.StatusOK, user, nil)
}

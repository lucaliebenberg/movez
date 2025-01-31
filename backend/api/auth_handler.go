package api

import "github.com/lucaliebenberg/movez/backend/types"

type AuthHandler struct {
	userStore db.UserStore
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthReponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

type genericResp struct {}

func NewAuthHandler(userStore db.UserStore)*AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

func (h *AuthHandler) HandleAuthenticate() {}

func invalidCredentials() {}


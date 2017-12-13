package auth

import (
	// Package conflict
	authCore "pawninn/domain/usecase/auth"
	"pawninn/domain/gateway/database"
	"net/http"
	"fmt"
)

type Auth struct {
	Prefix string
	Interactor authCore.AuthInteractor
	Controller Controller
}

func NewAuth(prefix string, database *database.Database) *Auth {
	auth := Auth{
		Prefix: prefix,
		Interactor: authCore.AuthInteractor{Database: database},
	}
	auth.Controller = Controller{Interactor: auth.Interactor}

	setUpRoutes(&auth)
	return &auth
}

func setUpRoutes(auth *Auth) {
	http.HandleFunc(fmt.Sprintf("%s/login", auth.Prefix), auth.Controller.handleLogin)
}
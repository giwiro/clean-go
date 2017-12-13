package auth

import (
	"net/http"
	"fmt"
	authCore "pawninn/domain/usecase/auth"
)

type Controller struct {
	Interactor authCore.AuthInteractor
}

func (authController *Controller) handleLogin(w http.ResponseWriter, r *http.Request) {
	resp, _ := authController.Interactor.Authenticate("giwirodavalos@gmail.com", "213")
	if resp != nil {
		fmt.Fprintf(w, "User found! -> %s", resp.Email)
	}else {
		fmt.Fprint(w, "Not found")
	}
}
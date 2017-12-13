package auth

import (
	"pawninn/domain/gateway/database"
	"pawninn/domain/entities"
)

type AuthInteractor struct {
	Database *database.Database
}

func (authinteractor *AuthInteractor) Authenticate(email string, password string) (user *entities.User, error string) {
	found := authinteractor.Database.User.GetUserByEmail(email)
	if found.Email != "" {
		user = &found
		return user, ""
	}else {
		return user, "not found"
	}
}

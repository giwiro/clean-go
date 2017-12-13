package model

import (
	// "database/sql"
	"pawninn/domain/entities"
	"pawninn/domain/gateway/database"
)

type UserGateway struct {
	database.UserGateway
	List map[string]entities.User
}

func NewUserGateway() *UserGateway {
	u := new(UserGateway)
	u.List = make(map[string]entities.User)
	u.List["giwirodavalos@gmail.com"] = entities.User{
		Id: 123,
		Email: "giwirodavalos@gmail.com",
		Password: "123456",
	}
	return u
}

func (userGateway *UserGateway) GetUserByEmail(email string) entities.User {
	return userGateway.List[email]
}
package model

import (
	"database/sql"
	"pawninn/domain/entities"
	"pawninn/domain/gateway/database"
)

type UserGateway struct {
	Connection *sql.DB
	database.UserGateway
}

func (userGateway *UserGateway) GetUserById(id uint32) entities.User {
	user := entities.User{Email: "giwirodavalos@gmail.com", Username: "giwiro"}
	return user
}
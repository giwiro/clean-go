package database

import "pawninn/domain/entities"

type Database struct {
	// Connection interface{}
	User interface{UserGateway}
}

type UserGateway interface {
	GetUserById(id uint32) entities.User
	GetUserByEmail(email string) entities.User
}


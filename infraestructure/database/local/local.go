package local

import (
	"pawninn/config"
	"pawninn/domain/gateway/database"
	"pawninn/infraestructure/database/local/model"
)

func NewLocalDatabase(config *config.Config) *database.Database {
	localDatabase := database.Database{
		User: model.NewUserGateway(),
	}
	return &localDatabase
}

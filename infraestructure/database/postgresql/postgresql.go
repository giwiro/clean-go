package postgresql

import (
	"database/sql"
	_"github.com/lib/pq"
	"pawninn/domain/gateway/database"
	"pawninn/infraestructure/database/postgresql/model"
	"fmt"
	"os"
	"pawninn/config"
)

func NewPostgresqlDatabase(config *config.Config) *database.Database {
	db, err := sql.Open("postgres", config.Postgresql.Uri)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	postgresqlDatabase := database.Database{
		User: model.UserGateway{Connection: db},
	}
	return &postgresqlDatabase
}

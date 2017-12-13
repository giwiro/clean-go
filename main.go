package main

import (
	// "pawninn/infraestructure/database/postgresql"
	"pawninn/infraestructure/database/local"
	"net/http"
	"log"
	"pawninn/delivery/http/auth"
	"pawninn/config"
	"fmt"
)

func main() {
	config := config.NewConfig()
	// database := postgresql.NewPostgresqlDatabase(config)
	database := local.NewLocalDatabase(config)
	auth.NewAuth("/auth", database)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}

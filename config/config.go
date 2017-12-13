package config

import (
	"os"
	"strconv"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

const (
	DefaultPostgresqlUri = "postgres://user:password@localhost/test"
	DefaultPort = 8080
	DefaultEnv = "dev"
)

type Config struct {
	Postgresql 		DatabaseConfig		`json:"postgresql"`
	Port 			uint64				`json:"port"`
	Environment 	string				`json:"environment"`
}

type DatabaseConfig struct {
	Uri 			string				`json:"uri"`
}

func readFile() *Config {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	c := new(Config)
	json.Unmarshal(raw, c)
	return c
}

func NewConfig() *Config {
	fileConfig := readFile()
	port, _ := strconv.ParseUint(getEnv("PORT", strconv.FormatUint(fileConfig.Port, 10),
		strconv.Itoa(DefaultPort)), 10, 32)
	config := Config{
		Postgresql: DatabaseConfig{
			Uri: getEnv("POSTGRESQL_URI", fileConfig.Postgresql.Uri, DefaultPostgresqlUri),
		},
		Environment: getEnv("ENV", "", DefaultEnv),
		Port: port,
	}
	return &config
}

// Order of read: envValue, fileValue, fallbackValue
func getEnv(key string, fileValue string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if fileValue != "" {
		return fileValue
	}
	return fallback
}

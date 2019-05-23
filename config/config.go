package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{}
}

func (databaseConfig *DatabaseConfig) NewPostgresSource() string {
	databaseConfig.Host = os.Getenv("POSTGRES_HOST")
	databaseConfig.Port = os.Getenv("POSTGRES_PORT")
	databaseConfig.User = os.Getenv("POSTGRES_USER")
	databaseConfig.DB = os.Getenv("POSTGRES_DB")
	databaseConfig.Password = os.Getenv("POSTGRES_PASSWORD")

	var dbSource string
	if databaseConfig.Password == "" {
		dbSource = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.DB)
	} else {
		dbSource = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.DB, databaseConfig.Password)
	}
	return dbSource
}

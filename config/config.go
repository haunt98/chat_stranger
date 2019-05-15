package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
}

func NewDatabaseConfig() *DatabaseConfig {
	var databaseConfig DatabaseConfig

	databaseConfig.Host = os.Getenv("DBHOST")
	databaseConfig.Port = os.Getenv("DBPORT")
	databaseConfig.User = os.Getenv("DBUSER")
	databaseConfig.Dbname = os.Getenv("DBNAME")
	databaseConfig.Password = os.Getenv("DBPASSWORD")

	return &databaseConfig
}

func (databaseConfig *DatabaseConfig) NewPostgresSource() string {
	var dbSource string
	if databaseConfig.Password == "" {
		dbSource = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Dbname)
	} else {
		dbSource = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Dbname, databaseConfig.Password)
	}
	return dbSource
}

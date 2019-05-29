package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Postgres struct{}

func NewPostgres() DatabaseReader {
	return &Postgres{}
}

func (postgres *Postgres) GetDBMS() string {
	return "postgres"
}

func (postgres *Postgres) GetSource() string {
	if viper.GetString("postgres.password") == "" {
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			viper.GetString("postgres.host"),
			viper.GetString("postgres.port"),
			viper.GetString("postgres.user"),
			viper.GetString("postgres.dbname"),
		)
	}

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.password"),
	)
}

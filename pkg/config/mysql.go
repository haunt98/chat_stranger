package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewMySQL() DatabaseReader {
	return &MySQL{}
}

type MySQL struct{}

func (mysql *MySQL) GetDBMS() string {
	return "mysql"
}

func (mysql *MySQL) GetSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"), viper.GetString("mysql.password"),
		viper.GetString("mysql.host"), viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
}

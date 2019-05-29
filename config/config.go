package config

import (
	"fmt"
	"github.com/1612180/chat_stranger/log"
	"github.com/1612180/chat_stranger/models"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DatabaseReader interface {
	GetDBMS() string
	GetSource() string
}

func CreateDefaultAdmin() *models.AdminUpload {
	viper.SetConfigName("cfg")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.ServerLog(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return &models.AdminUpload{
		Name:     viper.GetString("default_admin.name"),
		Password: viper.GetString("default_admin.password"),
		FullName: viper.GetString("default_admin.fullname"),
	}
}

func GetJWTSecretKey() string {
	viper.SetConfigName("cfg")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.ServerLog(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return viper.GetString("jwt.secret_key")
}

func GetPort() string {
	viper.SetConfigName("cfg")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.ServerLog(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return viper.GetString("port")
}

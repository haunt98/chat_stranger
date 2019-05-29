package config

import (
	"fmt"
	"github.com/1612180/chat_stranger/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DatabaseReader interface {
	GetDBMS() string
	GetSource() string
}

func Init() {
	viper.SetConfigName("cfg")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.ServerLog(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	viper.SetDefault("port", "8080")
}

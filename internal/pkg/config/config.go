package config

import (
	"github.com/1612180/chat_stranger/internal/pkg/variable"

	"github.com/spf13/viper"
)

type Config interface {
	Get(key string) string
	Set(key, value string)
}

func NewConfig(typeCfg string) Config {
	cfg := defaultConfig{
		m: make(map[string]string),
	}

	if typeCfg == variable.Viper {
		cfg.m[variable.Port] = viper.GetString(variable.Port)
		cfg.m[variable.DbDialect] = viper.GetString(variable.DbDialect)
		cfg.m[variable.DbUrl] = viper.GetString(variable.DbUrl)
		cfg.m[variable.DbMode] = viper.GetString(variable.DbMode)
		cfg.m[variable.JWTSecret] = viper.GetString(variable.JWTSecret)
		cfg.m[variable.GinMode] = viper.GetString(variable.GinMode)
		cfg.m[variable.Config] = typeCfg
	} else if typeCfg == variable.Test {
		cfg.m[variable.DbDialect] = "sqlite3"
		cfg.m[variable.DbUrl] = ":memory:"
		cfg.m[variable.DbMode] = "debug"
		cfg.m[variable.JWTSecret] = "secret"
		cfg.m[variable.GinMode] = "debug"
		cfg.m[variable.Config] = typeCfg
	}
	return &cfg
}

// implement

type defaultConfig struct {
	m map[string]string
}

func (c *defaultConfig) Get(key string) string {
	value, ok := c.m[key]
	if !ok {
		return ""
	}
	return value
}

func (c *defaultConfig) Set(key, value string) {
	c.m[key] = value
}

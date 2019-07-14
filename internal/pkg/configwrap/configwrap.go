package configwrap

import (
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/spf13/viper"
)

type Config interface {
	Get(key string) string
	Set(key, value string)
}

func NewConfig(mode string) Config {
	c := viperConfig{
		m: make(map[string]string),
	}

	if mode == "viper" {
		c.m[variable.Port] = viper.GetString(variable.Port)
		c.m[variable.DbDialect] = viper.GetString(variable.DbDialect)
		c.m[variable.DbUrl] = viper.GetString(variable.DbUrl)
		c.m[variable.DbMode] = viper.GetString(variable.DbMode)
		c.m[variable.JWTSecret] = viper.GetString(variable.JWTSecret)
		c.m[variable.GinMode] = viper.GetString(variable.GinMode)
	} else if mode == "test" {
		c.m[variable.DbDialect] = "sqlite3"
		c.m[variable.DbUrl] = ":memory:"
		c.m[variable.DbMode] = "debug"
		c.m[variable.JWTSecret] = "secret"
		c.m[variable.GinMode] = "debug"
	}

	return &c
}

// implement

type viperConfig struct {
	m map[string]string
}

func (c *viperConfig) Get(key string) string {
	value, ok := c.m[key]
	if !ok {
		return ""
	}
	return value
}

func (c *viperConfig) Set(key, value string) {
	c.m[key] = value
}

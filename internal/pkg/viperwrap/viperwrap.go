package viperwrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	ConfigUrl = "configurl"
)

type Environment struct {
	Name        string
	ServiceName string
	ProjectName string
	ConfigJson  string
}

type ConfigResponse struct {
	Code    int         `json:"returncode"`
	Message string      `json:"returnmessage"`
	Data    Environment `json:"data"`
}

func get(url string, data interface{}) bool {
	fields := logrus.Fields{
		"event":  "config",
		"where":  "remote",
		"action": "get",
	}

	// client with timeout
	client := http.Client{
		Timeout: time.Second * 30,
	}

	// get request
	res, err := client.Get(url)
	if err != nil {
		logrus.WithFields(fields).
			Error(err)
		return false
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			logrus.WithFields(fields).
				Error(err)
		}
	}()

	// read response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.WithFields(fields).
			Error(err)
		return false
	}

	// parse json
	if err := json.Unmarshal(body, data); err != nil {
		logrus.WithFields(fields).
			Error(err)
		return false
	}
	return true
}

func loadRemote(url string) ([]byte, bool) {
	fields := logrus.Fields{
		"event": "config",
		"where": "remote",
	}
	logrus.WithFields(fields).
		Info("Start loading remote config")

	if url == "" {
		logrus.WithFields(fields).
			Error("Missing remote url")
		return nil, false
	}

	configRes := ConfigResponse{}
	if ok := get(url, &configRes); !ok {
		return nil, false
	}

	if configRes.Code != 1 {
		logrus.WithFields(fields).
			Errorf("Load config from remote failed with code :%d", configRes.Code)
		return nil, false
	}
	return []byte(configRes.Data.ConfigJson), true
}

func loadLocal(serviceName, configFile, configPath string) {
	fields := logrus.Fields{
		"event": "config",
		"where": "local",
	}
	logrus.WithFields(fields).
		Info("Start loading local config")

	viper.SetConfigName(configFile)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", serviceName))
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithFields(fields).
			Error(err)
	}
}

func Load(serviceName, configFile, configPath string) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	remoteConfig, ok := loadRemote(viper.GetString(ConfigUrl))
	if !ok {
		loadLocal(serviceName, configFile, configPath)
		return
	}

	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewBuffer(remoteConfig)); err != nil {
		logrus.WithFields(logrus.Fields{
			"event": "config",
		}).Error(err)
		loadLocal(serviceName, configFile, configPath)
	}
}

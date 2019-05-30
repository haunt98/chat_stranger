package configutils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/1612180/chat_stranger/pkg/network"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var (
	CONFIG_URL_ENTRY          = "configurl"
	REMOTE_CONFIG_LOG_PREFIX  = "[REMOTE CONFIG]"
	LOCAL_CONFIG_LOG_PREFIX   = "[LOCAL CONFIG]"
	WARNING_LEVEL             = "[WARNING]"
	ERROR_LEVEL               = "[ERROR]"
	INFO_LEVEL                = "[INFO]"
	CONFIG_URL_REQUIRED_ERROR = errors.New("remote config url is required")
)

func logConfigMessage(prefix, level, msg string) {
	log.Println(fmt.Sprintf("%s %s %s", prefix, level, msg))
}

func LoadRemoteConfiguration(url string) ([]byte, error) {
	if len(url) == 0 {
		return nil, CONFIG_URL_REQUIRED_ERROR
	}

	httpClient := network.NewClient()
	ctx := context.Background()
	resp := ConfigServiceResponse{}
	_, err := httpClient.Get(ctx, url, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ReturnCode != 1 {
		return nil, fmt.Errorf("Loading configuration from remote failed with StatusCode :%d", resp.ReturnCode)
	}

	logConfigMessage(
		REMOTE_CONFIG_LOG_PREFIX,
		INFO_LEVEL,
		"Loaded configuration from remote successfully",
	)

	log.Printf(
		"[Configuration Info]: PrjectName: `%s`, ServiceName: `%s`, EnvName: `%s`",
		resp.Data.ProjectName,
		resp.Data.ServiceName,
		resp.Data.EnvName,
	)

	return []byte(resp.Data.ConfigJson), nil
}

func LoadConfiguration(serviceName, configFile, configPath string) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	configURL := viper.GetString(CONFIG_URL_ENTRY)
	configByte, err := LoadRemoteConfiguration(configURL)
	if err == nil {
		viper.SetConfigType("json")
		err = viper.ReadConfig(bytes.NewBuffer(configByte))
		if err == nil {
			logConfigMessage(
				REMOTE_CONFIG_LOG_PREFIX,
				INFO_LEVEL,
				"Read configuration from remote successfully",
			)
			return
		}
	}

	logConfigMessage(
		REMOTE_CONFIG_LOG_PREFIX,
		ERROR_LEVEL,
		err.Error(),
	)
	logConfigMessage(
		LOCAL_CONFIG_LOG_PREFIX,
		WARNING_LEVEL,
		"Loading configuration from local file system as a fallback",
	)
	viper.SetConfigName(configFile)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", serviceName))
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		logConfigMessage(
			LOCAL_CONFIG_LOG_PREFIX,
			ERROR_LEVEL,
			"Can't load configuration from local file system",
		)
	}

	logConfigMessage(
		LOCAL_CONFIG_LOG_PREFIX,
		INFO_LEVEL,
		"Loaded configuration from local file system successfully",
	)
}

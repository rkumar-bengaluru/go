package config

import (
	"errors"
	"fmt"

	"github.com/rkumar-bengaluru/go/logger"
	"github.com/spf13/viper"
)

type EcommerceConfigReader interface {
	Initialize() error

	GetBooleanKey(k string) (bool, error)
	GetIntegerKey(k string) (int, error)
	GetStringKey(k string) (string, error)
}

type EcommConfigReader struct {
	Logger  *logger.Logger
	Default *viper.Viper
}

func (c EcommConfigReader) Initialize() error {
	c.Logger.Info("initializing configuration....")
	channelDefault := make(chan error)

	go func() {
		c.Logger.Info(fmt.Sprintf("default config path %v", ConfigPath))
		channelDefault <- c.init(ConfigPath)
	}()

	errDefault := <-channelDefault
	if errDefault != nil {
		fmt.Println(errDefault.Error())
		return errDefault
	}

	return nil
}

func (c EcommConfigReader) GetBooleanKey(key string) (bool, error) {
	if c.Default.IsSet(key) {
		return c.Default.GetBool(key), nil
	}
	return false, nil
}

func (c EcommConfigReader) GetIntegerKey(key string) (int, error) {
	if c.Default.IsSet(key) {
		return c.Default.GetInt(key), nil
	}

	return 0, errors.New("coud not find " + key)
}

func (c EcommConfigReader) GetStringKey(key string) (string, error) {
	if c.Default.IsSet(key) {
		return c.Default.GetString(key), nil
	}
	return "", nil
}

func NewConfigReader(log *logger.Logger) *EcommConfigReader {

	config := &EcommConfigReader{
		Default: viper.New(),
		Logger:  log,
	}

	return config
}

func (c EcommConfigReader) init(path string) error {
	c.Default.AddConfigPath(path)
	c.Default.SetConfigType(ConfigType)
	c.Default.SetConfigName(ConfigName)

	r := c.Default.ReadInConfig()

	return r
}

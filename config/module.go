package config

import (
	"errors"
	"fmt"

	"github.com/rkumar-bengaluru/go/logger"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewConfigReader,
	),
	fx.Invoke(
		Initialize,
	),
)

var ConfigPath []string = []string{"./resources", "../resources", "."}
var ConfigType string = "yml"
var ConfigName string = "config"

type ConfigReaderI interface {
	GetBooleanKey(key string) (bool, error)
	GetIntegerKey(key string) (bool, error)
	GetStringKey(key string) (bool, error)
}

type ConfigReader struct {
	log     *logger.Logger
	Default *viper.Viper
}

func NewConfigReader(log *logger.Logger) *ConfigReader {
	return &ConfigReader{
		log:     log,
		Default: viper.New(),
	}
}

func (c *ConfigReader) GetBooleanKey(key string) (bool, error) {
	if c.Default.IsSet(key) {
		return c.Default.GetBool(key), nil
	}
	return false, errors.New("cannot found the key")
}

func (c *ConfigReader) GetIntegerKey(key string) (int, error) {
	c.log.Info(fmt.Sprintf("is key set %v\n", c.Default.IsSet(key)))
	if c.Default.IsSet(key) {
		c.log.Info(fmt.Sprintf("found key %v value %v\n", key, c.Default.GetInt(key)))
		return c.Default.GetInt(key), nil
	}
	return 0, errors.New("cannot found the key " + key)
}

func (c *ConfigReader) GetStringKey(key string) (string, error) {
	if c.Default.IsSet(key) {
		return c.Default.GetString(key), nil
	}
	return "", errors.New("cannot found the key")
}

func Initialize(c *ConfigReader) {
	c.log.Info("intializing config reader...")
	for _, path := range ConfigPath {
		c.log.Info(fmt.Sprintf("adding config path %v\n", path))
		c.Default.AddConfigPath(path)
	}
	c.Default.SetConfigType(ConfigType)
	c.Default.SetConfigName(ConfigName)
	err := c.Default.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	c.log.Info("started config successfully...")
}

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

var (
	ConfigPath []string = []string{"./resources", "../resources", "."}
	ConfigType string   = "yml"
	ConfigName string   = "config"
)

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

func (c *NewConfigReader) GetBooleanKey(key string) (bool, error) {
	if c.Default.IsSet(key) {
		return c.Default.GetBool(key), nil
	}
	return false, errors.New("cannot found the key")
}

func (c *NewConfigReader) GetIntegerKey(key string) (int, error) {
	if c.Default.IsSet(key) {
		return c.Default.GetInt(key), nil
	}
	return 0, errors.New("cannot found the key")
}

func (c *NewConfigReader) GetStringKey(key string) (int, error) {
	if c.Default.IsSet(key) {
		return c.Default.GetString(key), nil
	}
	return 0, errors.New("cannot found the key")
}

func Initialize(c *NewConfigReader) {
	c.log.Info("intializing config reader...")
	for path := range ConfigPath {
		c.log.Info(fmt.Sprintf("adding config path %v\n", path))
		c.Default.AddConfigPath(path)
	}
	c.Default.ConfigPath(ConfigType)
	c.Default.ConfigName(ConfigName)
	err := c.Default.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

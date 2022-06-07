package configsreader

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ConfigsReader interface {
	Init() error

	GetBool(key string) bool
	GetInt(key string) int
	GetProfile() string
	GetString(key string) string
}

type ViperReader struct {
	profileName string
	Default     *viper.Viper
	Profile     *viper.Viper
}

func (c ViperReader) Init() error {
	channelDefault := make(chan error)
	channelProfile := make(chan error)

	go func() {
		channelDefault <- c.initDefault(ConfigPath)
	}()

	go func() {
		channelProfile <- c.initProfile(ConfigPath)
	}()

	errDefault := <-channelDefault
	if errDefault != nil {
		return errDefault
	}

	errProfile := <-channelProfile
	if errProfile != nil {
		return errProfile
	}

	return nil
}

func (c ViperReader) GetBool(key string) bool {
	if c.Profile.IsSet(key) {
		return c.Profile.GetBool(key)
	}

	return c.Default.GetBool(key)
}

func (c ViperReader) GetInt(key string) int {
	if c.Profile.IsSet(key) {
		return c.Profile.GetInt(key)
	}

	return c.Default.GetInt(key)
}

func (c ViperReader) GetString(key string) string {
	if c.Profile.IsSet(key) {
		return c.Profile.GetString(key)
	}

	return c.Default.GetString(key)
}

func (c ViperReader) GetProfile() string {
	return c.profileName
}

func newViperReader() *ViperReader {
	profileName := os.Getenv(ConfigProfile)
	if profileName == "" {
		profileName = ConfigProfileDefault
	}

	config := &ViperReader{
		profileName: profileName,
		Default:     viper.New(),
		Profile:     viper.New(),
	}

	return config
}

func (c ViperReader) initDefault(path string) error {
	c.Default.AddConfigPath(path)
	c.Default.SetConfigType(ConfigType)
	c.Default.SetConfigName(ConfigName)

	return c.Default.ReadInConfig()
}

func (c ViperReader) initProfile(path string) error {
	c.Profile.AddConfigPath(path)
	c.Profile.SetConfigType(ConfigType)
	c.Profile.SetConfigName(fmt.Sprintf("%s-%s", ConfigName, c.profileName))
	c.Profile.SetEnvPrefix("")
	c.Profile.AutomaticEnv()
	c.Profile.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", ""))

	return c.Profile.ReadInConfig()
}

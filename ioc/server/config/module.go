package configs

import (
	"go.uber.org/fx"

	reader "github.com/rkumar-bengaluru/go/ioc/server/config/configreader"
)

var Module = fx.Options(
	fx.Provide(
		newConfig,
	),
	fx.Invoke(
		loadConfig,
	),
)

type in struct {
	fx.In
	Server *ServerConfig,
}

type out struct {
	fx.Out

	Server *ServerConfig,
}

func newConfig() out {
	return out{
		Server: &ServerConfig{},
	}
}

func loadConfig(
	configReader reader.ConfigsReader,
	config in,
) {
	config.Server.Load(configReader)
}

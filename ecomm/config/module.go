package config

import (
	"github.com/rkumar-bengaluru/go/logger/console"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newConfig,
	),
	fx.Invoke(
		initConfigReader,
	),
)

func newConfig(log *console.ConsoleLogger) *EcommConfigReader {
	return NewConfigReader(log)
}

func initConfigReader(c *EcommConfigReader) error {
	err := c.Initialize()
	return err
}

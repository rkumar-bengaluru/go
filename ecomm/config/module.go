package config

import (
	"github.com/rkumar-bengaluru/go/ecomm/log"
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

func newConfig(log *log.DevelopmentLogger) *EcommConfigReader {
	return NewConfigReader(log)
}

func initConfigReader(c *EcommConfigReader) error {
	err := c.Initialize()
	return err
}

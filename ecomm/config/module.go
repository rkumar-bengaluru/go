package config

import (
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

func newConfig() *EcommConfigReader {
	return NewConfigReader()
}

func initConfigReader(c EcommerceConfigReader) error {
	err := c.Initialize()
	return err
}

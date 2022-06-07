package configsreader

import (
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newReader,
	),
	fx.Invoke(
		initConfigReader,
	),
)

type out struct {
	fx.Out

	ConfigsReader ConfigsReader
}

func newReader() out {
	return out{
		ConfigsReader: newViperReader(),
	}
}

func initConfigReader(configsReader ConfigsReader) error {
	err := configsReader.Init()
	fmt.Printf("Profile: %s\n", configsReader.GetProfile())

	return err
}

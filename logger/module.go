package logger

import (
	"github.com/rkumar-bengaluru/go/logger/console"
	"github.com/rkumar-bengaluru/go/logger/file"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(console.Module, file.Module),
)

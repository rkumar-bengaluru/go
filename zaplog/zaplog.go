package zaplog

import (
	"encoding/json"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func Initialize() {
	logCfg := []byte(`{
			"level" : "debug",
			"encoding" : "json",
			"outputPaths" : "stdout",
			"errorOutputPaths" : "stderr",
			"encoderConfig" : {
				"messageKey" : "message",
				"levelKey" : "level",
				"levelEncoder" : "lowercase"
			}
		}`)

	var cfg zap.Config
	if err := json.Unmarshal(logCfg, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
}

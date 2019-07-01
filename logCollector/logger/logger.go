package logger

import (
"github.com/astaxie/beego/logs"
"logCollector/config"
)

func caseLogLevel(level string) (logLevel int) {
	switch level {
	case "debug":
		logLevel = logs.LevelDebug
	case "info":
		logLevel = logs.LevelInfo
	case "warn":
		logLevel = logs.LevelWarn
	case "error":
		logLevel = logs.LevelError
	default:
		logLevel = logs.LevelInfo
	}
	return logLevel
}

func InitLogger() {
	logs.SetLevel(caseLogLevel(config.LogLevel))
	logs.SetLogger("console")
	logs.SetLogger("file", `{"filename":"logs/logCollector.log","color":true}`)
}

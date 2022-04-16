package constant

import "errors"

type LogLevel string

var (
	ErrInvalidLogLevel = errors.New("invalid log level")
)

const (
	LogLevelInfo    = LogLevel("info")
	LogLevelError   = LogLevel("error")
	LogLevelWarn    = LogLevel("warn")
	LogLevelUnknown = LogLevel("unknown")
)

var logLevelMap = map[string]LogLevel{
	"info":    LogLevelInfo,
	"error":   LogLevelError,
	"warn":    LogLevelWarn,
	"warning": LogLevelWarn,
}

func GetLogLevel(value string) (logLevel LogLevel, err error) {
	logLevel, ok := logLevelMap[value]
	if !ok {
		err = ErrInvalidLogLevel
		logLevel = LogLevelUnknown
	}
	return
}

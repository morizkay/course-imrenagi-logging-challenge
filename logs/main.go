package logs

import (
	"os"
	"runtime/debug"
	"strings"

	"github.com/rs/zerolog"
)

func NewLogger(config LoggerConfig) (*Logger, error) {
	level, err := parseLevel(config.Level)
	if err != nil {
		return nil, err
	}

	if config.TimeFormat != "" {
		zerolog.TimeFieldFormat = config.TimeFormat
	}

	if config.StackTrace {
		zerolog.ErrorStackMarshaler = func(err error) interface{} {
			return string(debug.Stack())
		}
	}

	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()

	if config.Caller {
		logger = logger.With().Caller().Logger()
	}

	return &Logger{logger}, nil
}

func parseLevel(level string) (zerolog.Level, error) {
	switch strings.ToLower(level) {
	case "trace":
		return zerolog.TraceLevel, nil
	case "debug":
		return zerolog.DebugLevel, nil
	case "info", "":
		return zerolog.InfoLevel, nil
	case "warn", "warning":
		return zerolog.WarnLevel, nil
	case "error":
		return zerolog.ErrorLevel, nil
	case "fatal":
		return zerolog.FatalLevel, nil
	case "panic":
		return zerolog.PanicLevel, nil
	default:
		return zerolog.InfoLevel, nil
	}
}

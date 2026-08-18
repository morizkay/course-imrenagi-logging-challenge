package logs

import (
	"io"
	"os"
	"path/filepath"
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

	writer, err := buildWriter(config)
	if err != nil {
		return nil, err
	}

	logger := zerolog.New(writer).Level(level).With().Timestamp().Logger()

	if config.Caller {
		logger = logger.With().Caller().Logger()
	}

	return &Logger{logger}, nil
}

func buildWriter(config LoggerConfig) (io.Writer, error) {
	output := strings.ToLower(config.Output)
	if output == "" {
		output = "stdout"
	}

	switch output {
	case "stdout":
		return os.Stdout, nil
	case "file":
		return openLogFile(config.FilePath)
	case "both":
		lf, err := openLogFile(config.FilePath)
		if err != nil {
			return nil, err
		}
		return zerolog.MultiLevelWriter(os.Stdout, lf), nil
	default:
		return os.Stdout, nil
	}
}

func openLogFile(path string) (*os.File, error) {
	if path == "" {
		path = "logs/app.log"
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
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

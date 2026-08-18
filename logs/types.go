package logs

import "github.com/rs/zerolog"

type LoggerConfig struct {
	Level      string
	TimeFormat string
	Caller     bool
	StackTrace bool
	// Output controls log destination: "stdout", "file", or "both".
	Output string
	// FilePath is the log file path when Output is "file" or "both".
	FilePath string
}

type Logger struct {
	zerolog.Logger
}

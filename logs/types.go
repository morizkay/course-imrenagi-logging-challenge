package logs

import "github.com/rs/zerolog"

type LoggerConfig struct {
	Level      string
	TimeFormat string
	Caller     bool
	StackTrace bool
}

type Logger struct {
	zerolog.Logger
}

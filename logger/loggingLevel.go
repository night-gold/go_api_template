package logger

import (
	"os"
	"github.com/rs/zerolog"
)

// LoggingLevel start the infos logger.
func LoggingLevel(lvl string) zerolog.Logger{
	var infos zerolog.Logger

	switch lvl {
		case "debug":
			infos = zerolog.New(os.Stdout).With().Timestamp().Logger()
		case "infos":
			infos = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	return infos
}
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
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "infos":
			infos = zerolog.New(os.Stdout).With().Timestamp().Logger()
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
	}

	return infos
}
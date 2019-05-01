package api

import (
	"os"
	"fmt"

	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	"github.com/night-gold/go_api_template/logger"
)

func init(){
	//var infos zerolog.Logger
	errors := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Loglevel contains the LogLevel defined in the env variable, it will define the global loglevel and the log variables associated to output.
	// panic, fatal, error, warn, info, debug
	loglevel := os.Getenv("LOGLEVEL")

	infos := logger.LoggingLevel(loglevel)

	switch loglevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			infos.Info().Msg("Server is starting.")
		case "infos":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			infos.Info().Msg("Server is starting.")
		case "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
	}
	
	errors.Error().Msg("What an error") //declare message in error for starting http server
}

func main(){
	fmt.Println("test")
}
package main

import (
	"os"
	"fmt"

	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	"github.com/night-gold/go_api_template/logger"
)

var infos zerolog.Logger
var errors zerolog.Logger

func init(){
 	//var infos zerolog.Logger
	errors = zerolog.New(os.Stderr).With().Timestamp().Logger()

	/* Loglevel contains the LogLevel defined in the env variable, it will define the global loglevel and the log variables associated to output.
	panic, fatal, error, warn, info, debug */
	loglevel, ok := os.LookupEnv("LOGLEVEL")
	if !ok {
		loglevel = "infos"
	}

	infos = logger.LoggingLevel(loglevel)
	infos.Info().Msg("The serveur is starting.")
}

func main(){
	errors.Error().Msg("What an error")//declare message in error for starting http server
	fmt.Println("test")
}
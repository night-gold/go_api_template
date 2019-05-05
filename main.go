package main

import (
	"os"

	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	"github.com/night-gold/go_api_template/logger"

	"github.com/night-gold/go_api_template/utils"

	"github.com/night-gold/go_api_template/server"
)

var infos zerolog.Logger
var errors zerolog.Logger

func init(){
 	//Setting warning and more logger
	errors = zerolog.New(os.Stderr).With().Timestamp().Logger()

	/* Loglevel contains the LogLevel defined in the env variable, it will define the global loglevel and the log variables associated to output.
	panic, fatal, error, warn, info, debug */
	loglevel := utils.GetEnv("LOGLEVEL","infos")

	//Setting logging global level
	infos = logger.LoggingLevel(loglevel)
	infos.Info().Msg("The serveur is starting.")
}

func main(){
	r := server.NewServer(infos, errors)
}
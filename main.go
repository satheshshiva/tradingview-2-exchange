package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/satheshshiva/go-banner-printer/banner"
	handlers "github.com/satheshshiva/tradingview-2-exchange/handler"

	"net/http"
	"os"
)

func main() {
	//server port
	serverPort := port
	// if port env found then use it
	if p, ok := os.LookupEnv("PORT"); ok {
		serverPort = p
	}

	//print banner
	_ = banner.Print(nil)

	//initilize loggers
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	//initialize handler
	handlers.RegisterRootEndpoint("/")

	//start the server
	log.Info().Msgf("Starting server on port: %v", serverPort)
	log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)).Msg("")
}

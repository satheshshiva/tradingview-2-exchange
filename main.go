package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/satheshshiva/go-banner-printer/banner"
	handlers "github.com/satheshshiva/tradingview-2-exchange/services"
	"github.com/satheshshiva/tradingview-2-exchange/services/tradingview"

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

	// init logger
	initLogger()

	//initialize handler
	handlers.RegisterRootEndpoint("/")
	tv := tradingview.New()
	if err := tv.RegisterHandler("/tradingview"); err != nil {
		log.Fatal().Msgf("error occured while registering tradingview handler: %s", err)
	}

	//start the server
	log.Info().Msgf("Starting server on port: %v", serverPort)
	log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)).Msg("")
}

func initLogger() {
	if debugLogging {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	//initilize loggers
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
}

package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/satheshshiva/go-banner-printer/banner"
	handlers "github.com/satheshshiva/tradingview-2-exchange/services"
	"github.com/satheshshiva/tradingview-2-exchange/services/tradingview"
	"github.com/satheshshiva/tradingview-2-exchange/util"

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
	//beautify logger for local. Output only json logging in production
	if !util.IsProd() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	log.Logger = log.With().Caller().Logger()
	// to easily search the manual logs in GCP. use `jsonPayload.appname:"tradingview-2-exchange"`
	log.Logger = log.With().Str("appname", "tradingview-2-exchange").Logger()
}

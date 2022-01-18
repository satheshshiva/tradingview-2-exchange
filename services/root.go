package handlers

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func RegisterRootEndpoint(subUrl string) {
	http.HandleFunc(subUrl, func(w http.ResponseWriter, r *http.Request) {
		profile := struct {
			AppName string
			Uptime  string
			Healthy bool
		}{"Tradingview-2-Exchange", time.Since(startTime).String(), true}

		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(js)
		if err != nil {
			log.Print("Error writing to response", err)
		}
	})
}

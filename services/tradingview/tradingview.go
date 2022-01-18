package tradingview

import (
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

type TradingView struct {
}

func New() *TradingView {
	return &TradingView{}
}

func (tv *TradingView) RegisterHandler(subUrl string) error {
	http.HandleFunc(subUrl, tv.handle())
	return nil
}

func (tv *TradingView) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid Method", http.StatusInternalServerError)
			return
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Err(err).Msg("Error reading tradingview request body")
		}
		log.Debug().Msgf("Received TradingView Request: %s", b)

	}
}

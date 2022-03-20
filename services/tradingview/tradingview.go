package tradingview

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/satheshshiva/tradingview-2-exchange/services/exchange"
	"github.com/satheshshiva/tradingview-2-exchange/services/twitter"
	"github.com/satheshshiva/tradingview-2-exchange/util"
)

type TradingView struct {
	ex			exchange.Exchange
	passphrase	string
	twtr		*twitter.Twitter
}

func New(ex exchange.Exchange, passphrase string, twtr *twitter.Twitter) *TradingView {
	return &TradingView{ex: ex, passphrase: passphrase, twtr: twtr}
}

func (tv *TradingView) RegisterHandler(subUrl string) error {
	http.HandleFunc(subUrl, tv.handle())
	return nil
}

func (tv *TradingView) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check input method
		if r.Method != "POST" {
			http.Error(w, "Invalid Method", http.StatusInternalServerError)
			return
		}

		// read the input data
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Err(err).Msg("Error reading tradingview request body")
		}
		log.Info().Msgf("Received TradingView Request: %s", util.RedactJsonField(string(b), "passphrase"))

		// unmarshal input data
		p := &Payload{}
		if err = json.Unmarshal(b, p); err != nil {
			http.Error(w, "Failed unmarshaling data", http.StatusInternalServerError)
			log.Err(err).Msgf("Error unmarshaling tradingview request:%s", err.Error())
			return
		}

		// validate input
		if !tv.validateInput(p) {
			http.Error(w, "Invalid input parameters", http.StatusBadRequest)
			return
		}

		go func ()  {
			//call the exchange
		err = tv.ex.Trade(&exchange.NewTrade{
			Symbol: p.Ticker,
			Side:   p.Side,
			Type:   p.Type,
			Qty:    p.Size,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		}()

		if p.PostTweet && tv.twtr!=nil{
			go func ()  {
				tv.twtr.TweetTrade(p.Side, p.Comment)
			}()
		}
		
	}
}

func (tv *TradingView) validateInput(p *Payload) bool {
	if p.Side == "" || (strings.ToLower(p.Side) != "buy" && strings.ToLower(p.Side) != "sell") {
		log.Error().Msgf("Side is not `buy` or `sell`")
		return false
	}
	if p.Passphrase == "" || p.Passphrase != tv.passphrase {
		log.Error().Msgf("Tradingview passphrase mismatch or empty")
		return false
	}
	return true
}

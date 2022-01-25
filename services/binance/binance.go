package binance

import (
	"github.com/google/go-querystring/query"
	"github.com/rs/zerolog/log"
	"github.com/satheshshiva/tradingview-2-exchange/services/exchange"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

const (
	prodUrl          = "https://fapi.binance.com"
	testnetUrl       = "https://testnet.binancefuture.com"
	newTradeEndpoint = "/fapi/v1/order"
)

func New(apiKey string, apiSecret string, prodEnv bool) *binance {
	url := ""
	if prodEnv {
		url = prodUrl
	} else {
		url = testnetUrl
	}
	log.Info().Msgf("Initializing binance rest api with url %s", url)
	return &binance{url: url, apiKey: apiKey, apiSecret: apiSecret}
}

func (b *binance) Trade(n *exchange.NewTrade) error {
	nt := newTradeRequest{
		Symbol:    n.Symbol,
		Side:      n.Side,
		OrderType: n.Type,
		Qty:       n.Qty,
		Timestamp: "",
		Signature: "",
	}

	var v url2.Values
	var err error
	if v, err = query.Values(nt); err != nil {
		log.Err(err).Msgf("error while changing struct to query params")
		return err
	}
	url := b.url + newTradeEndpoint + "?" + v.Encode()
	if resp, err := http.Get(url); err != nil {
		log.Err(err).Msgf("error response from binance new trade request api end point")
		return err
	} else {
		respStr, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode == http.StatusOK {
			log.Info().Msgf("Response from binance new trade api HTTP:%v:%s", resp.StatusCode, respStr)
		} else {
			log.Error().Msgf("Response from binance new trade api HTTP:%v:%s", resp.StatusCode, respStr)
		}
	}
	return nil
}

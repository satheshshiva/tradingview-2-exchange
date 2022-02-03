package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/rs/zerolog/log"
	"github.com/satheshshiva/tradingview-2-exchange/services/exchange"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
	"time"
)

const (
	prodUrl          = "https://fapi.binance.com"
	testnetUrl       = "https://testnet.binancefuture.com"
	newTradeEndpoint = "/fapi/v1/order"
	headerApiKey     = "X-MBX-APIKEY"
)

func New(apiKey string, apiSecret string, prodEnv bool) *binance {
	url := ""
	if prodEnv {
		url = prodUrl
	} else {
		url = testnetUrl
	}
	log.Trace().Msgf("Initializing binance rest api with url %s", url)
	return &binance{url: url, apiKey: apiKey, apiSecret: apiSecret}
}

func (b *binance) Trade(n *exchange.NewTrade) error {
	timestamp := time.Now()
	nt := newTradeRequest{
		Symbol:    n.Symbol,
		Side:      strings.ToUpper(n.Side),
		OrderType: n.Type,
		Qty:       n.Qty,
		Timestamp: strconv.FormatInt(timestamp.UnixNano()/int64(time.Millisecond), 10),
	}

	var v url2.Values
	var err error
	if v, err = query.Values(nt); err != nil {
		log.Err(err).Msgf("error while changing struct to query params")
		return err
	}
	qp := v.Encode()
	url := b.url + newTradeEndpoint + "?" + qp + "&signature=" + b.signature(qp)
	log.Debug().Msg(url)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Err(err).Msg("error creating http request")
		return err
	}
	req.Header.Set(headerApiKey, b.apiKey)
	client := &http.Client{}
	if resp, err := client.Do(req); err != nil {
		log.Err(err).Msgf(err.Error())
		return err
	} else {
		respStr, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode == http.StatusOK {
			log.Info().Msgf("Successful response from binance new trade api HTTP:%v:%s", resp.StatusCode, respStr)
		} else {
			err = errors.New(fmt.Sprintf("Response from binance new trade api HTTP:%v:%s", resp.StatusCode, respStr))
			log.Err(err).Msg(err.Error())
			return err
		}
	}
	return nil
}

func (b *binance) GetApiUrl() string {
	return b.url
}

func (b *binance) signature(qp string) string {
	h := hmac.New(sha256.New, []byte(b.apiSecret))
	h.Write([]byte(qp))
	return hex.EncodeToString(h.Sum(nil))
}

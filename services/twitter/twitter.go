package twitter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/dghubble/oauth1"
	"github.com/rs/zerolog/log"
)

type Twitter struct {
	ApiKey 	string
	ApiSecret  string
	AccessToken 	string
	TokenSecret		string
}

const(
	postTweetEndpoint = "https://api.twitter.com/2/tweets"
)


func New(apiKey string, apiSecret string, accessToken string, tokenSecret string) *Twitter {
	return &Twitter{ApiKey: apiKey, ApiSecret: apiSecret, AccessToken: accessToken, TokenSecret: tokenSecret}
}

func (t *Twitter)TweetTrade(side string, comment string) {
	msg := "\n\n🤖🤖 Automated ** Scalping Bot ** Trade Alert\n\n"
	msg += "#Bitcoin\n\n"
	switch strings.ToLower(side){
		case strings.ToLower("buy"):
			msg += "LONG 🔼\n\n"
		case strings.ToLower("sell"):
			msg += "SHORT 🔻\n\n"
	}
	msg+=comment 
	t.tweet(msg)
}

func (t *Twitter)tweet(msg string) {
	config := oauth1.NewConfig(t.ApiKey, t.ApiSecret)
    token := oauth1.NewToken(t.AccessToken, t.TokenSecret)
    httpClient := config.Client(oauth1.NoContext, token)
	values := map[string]string{"text": msg}
    json_data, err := json.Marshal(values)
    if err != nil {
		log.Err(err).Msg("Error while marshaling Twitter API request")
    }
    resp, err := httpClient.Post(postTweetEndpoint, "application/json", bytes.NewBuffer(json_data ))
	if err != nil {
		log.Err(err).Msg("Error while calling Twitter API")
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("Error while decoding response from Twitter API")
    }

	log.Info().Msgf("Response from Twitter API: %s", string(body))

}
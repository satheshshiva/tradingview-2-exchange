package tradingview

type Payload struct {
	Strategy   string  `json:"strategy"` // name of the strategy
	Side       string  `json:"side"`     // side buy or sell
	Size       float32 `json:"size"`
	Time       string  `json:"time"`
	Ticker     string  `json:"ticker"`
	Exchange   string  `json:"exchange"` // exchange in which the chart was made
	Close      float32 `json:"close"`    // close price
	Passphrase string  `json:"passphrase"`
	Comment		string	 `json:"comment"`
	PostTweet	bool	 `json:"postTweet"`
	Type       string  `json:"type"`
}

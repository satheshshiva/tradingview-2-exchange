package exchange

type Exchange interface {
	Trade(n *NewTrade) error
}

type NewTrade struct {
	Symbol string
	Side   string
	Type   string
	Qty    float32
}

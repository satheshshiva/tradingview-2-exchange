package exchange

import "net/http"

type Exchange interface {
	Trade(n *NewTrade) error
	GetApiUrl() string
}

type NewTrade struct {
	Symbol string
	Side   string
	Type   string
	Qty    float32
}

func RegisterGetApiUrlHandler(ex Exchange, url string) error {
	http.HandleFunc(url, handle(ex))
	return nil
}

func handle(ex Exchange) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(ex.GetApiUrl()))
	}
}

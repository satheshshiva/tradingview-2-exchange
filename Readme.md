# TradingView-2-Exchange
This application is used to place a buy/sell order in Binance when your TradingView strategy triggers a long/short entry.

## Environmental Variables
1. BINANCE_API_KEY - Api key
2. BINANCE_API_SECRET - Api secret
3. BINANCE_PRODUCTION - Points the binance to production url or test net url. Possible values are `true` or `false`. Defaults to testnet.
4. TV_PASSPHRASE - The incoming trading view request payload should have this password under the `passphrase` field
5. PRODUCTION - Dont set for test env
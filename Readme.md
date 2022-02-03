# TradingView-2-Exchange
[TradingView](https://www.tradingview.com) provided an awesome feature to create a strategy in pinescript and backtest the strategy. The missing functionality in trading view is the lack of automatically placing the buy/sell orders in your exchange account, so that your strategy from tradingview can act like a tradingbot. This repo fixes the gap. This app acts as a bridge from your tradingview strategy to your exchange account.

### How its done?
This app takes advantage of the Alarm alert for the strategy in tradingview. To enable automatic trading in the tradingview, enable alarm and then enable the "Webhook URL" in the alert window. Deploy this application to cloud services like GCP(Google Cloud Console)/AWS or any other services and provide the application URL in the Webhook URL field. Replace the Message field with [trading-view-alert.json](sample/trading-view-alert.json). Whenever tradingview raises an alarm alert, this application will be called and buy/sell orders will be placed in your exchange account.

### Prerequisites
1. Go 1.16.
At this time GCP(Google Cloud Platform) supports only 1.16
2. Google Cloud Console account. Can use any other cloud based deployment services as well.
3. Binance Exchange account with API access enabled.

### Supported exchanges
1. Binance - complete support
2. Kucoin - tentative

### Environmental Variables
1. BINANCE_API_KEY - Api key
2. BINANCE_API_SECRET - Api secret
3. BINANCE_PRODUCTION - Points the binance to production url or test net url. Possible values are `true` or `false`. Defaults to testnet.
4. TV_PASSPHRASE - The incoming trading view request payload should have this password under the `passphrase` field
5. PRODUCTION - Dont set for test env

### Samples
The samples folder contains,
1. [app.yaml](sample/app.yaml) - For google cloud console deployment
2. [strategy.pinescript](strategy.pinescript) - a sample stragegy written in pine script. The script has a 70% success rate on BTCUSDT
3. [tradingview-alert.json](tradingview-alert.json) -  a sample tradingview webhook alert


### Note
You need to manually trigger the first trade by looking at the strategy and seeing what trade is in progress right now. Or if the first trade is already triggered by the bot, then you need to manually reduce the quantity by half. This is because of the one-cancels-other trading view strategy in tradingview.
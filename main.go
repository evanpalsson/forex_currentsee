package main

import (
"flag"
"fmt"

"github.com/santegoeds/oanda"
//"sort"
)

var (
	token   = flag.String("token", "416a2fa520421ce7561661dc35b4116b-27dce4b35b075d0223334ee89be7a02b", "Oanda authorization token.")
	account = flag.Int64("account", 8108490, "Oanda account.")
)

func main() {
	flag.Parse()

	if *token == "" {
		panic("An Oanda authorization token is required")
	}

	if *account == 0 {
		panic("An Oanda account is required")
	}

	client, err := oanda.NewFxPracticeClient(*token)
	if err != nil {
		panic(err)
	}

	client.SelectAccount(oanda.Id(*account))

	// List available instruments
	//instruments, err := client.Instruments(nil, nil)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(instruments)

	// Buy one unit of EUR/USD with a trailing stop of 10 pips.
	//tradeInfo, err := client.NewTrade(oanda.Buy, 1, "eur_usd", oanda.TrailingStop(10.0))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(tradeInfo)

	// Create and run a price server.
	//priceServer, err := client.NewPriceServer("eur_usd")
	//if err != nil {
	//	panic(err)
	//}
	//priceServer.ConnectAndHandle(func(instrument string, tick oanda.PriceTick) {
	//	fmt.Println("Received tick:", instrument, tick)
	//	priceServer.Stop()
	//})

	// Close the previously opened trade.
	//tradeCloseInfo, err := client.CloseTrade(tradeInfo.TradeId)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(tradeCloseInfo)

	candles, err := client.PollHighAskCandles("eur_usd", "H4")
	if err != nil {
		panic(err)
	}

	fmt.Println(candles)
	//fmt.Println(high)

}

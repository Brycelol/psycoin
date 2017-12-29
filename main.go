package main

import (
	"github.com/brycelol/psycoin/gdax"
)

func main() {
	priceTicker := gdax.BTCPriceTicker{PriceChannel: make(chan gdax.PriceTick, 60)}
	priceTicker.Start()
}

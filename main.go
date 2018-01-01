package main

import (
	"net/http"
	"time"
	"github.com/preichenberger/go-gdax"
	"os"
	"github.com/brycelol/psycoin/ticker"
	"github.com/brycelol/psycoin/trader"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {

	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	passphrase := os.Getenv("COINBASE_PASSPHRASE")

	client := gdax.NewClient(secret, key, passphrase)
	client.HttpClient = &http.Client{
		Timeout: 15 * time.Second,
	}

	tickChannel := make(chan ticker.PriceTick, 60)

	priceTicker := ticker.BTCPriceTicker{TickChannel: tickChannel}
	waitGroup.Add(1)
	go priceTicker.Start()

	waitGroup.Add(1)
	go trader.Trade(tickChannel)

	waitGroup.Wait()
}

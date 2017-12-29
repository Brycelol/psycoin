package gdax

import (
	"github.com/gorilla/websocket"
	"time"
	"fmt"
)

// Interface for multiple different price tickers.
type PriceTicker interface {
	Start()
	Stop()
}

type PriceTick struct {
	price     float64
	timestamp int64
}

// Concrete BTC Price Ticker
type BTCPriceTicker struct {
	PriceChannel chan PriceTick
}

var running = false

func (ticker *BTCPriceTicker) Start() {
	println("Starting BTCPriceTicker...")

	// Create WebSocket client.
	var wsDialer websocket.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil)

	if err != nil {
		println(err.Error())
	}

	// Create subscription message.
	subscribe := Message{
		Type: "subscribe",
		Channels: []MessageChannel{
			MessageChannel{
				Name: "ticker",
				ProductIds: []string{
					"BTC-GBP",
				},
			},
		},
	}

	// Subscribe to the price ticker
	if err := wsConn.WriteJSON(subscribe); err != nil {
		println(err.Error())
	}

	running = true

	message := Message{}
	for running {
		// Process messages spooled from the socket in JSON.
		if err := wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}

		if message.Type == "ticker" {
			// Write the price to channel.
			tick := PriceTick{price: message.Price, timestamp: time.Now().Unix()}
			ticker.PriceChannel <- tick

			// We have no consumers so we must read otherwise we would exhaust the channel.
			cTick := <-ticker.PriceChannel
			fmt.Printf("BTC Price £%f at %s\n", cTick.price, time.Unix(cTick.timestamp, 0))
		}
	}
}

func (ticker BTCPriceTicker) Stop() {
	println("Shutting down BTCPriceTicker...")
	running = false
}

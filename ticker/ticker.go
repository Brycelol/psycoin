package ticker

import (
	"github.com/preichenberger/go-gdax"
	"github.com/gorilla/websocket"
	"time"
)

// Interface for multiple different price tickers.
type PriceTicker interface {
	Start()
	Stop()
}

type PriceTick struct {
	Price float64
	High  float64
	Time  time.Time
}

// Concrete BTC Price Ticker
type BTCPriceTicker struct {
	TickChannel chan PriceTick
}

var running = false

func (ticker *BTCPriceTicker) Start() {
	println("[BTCPriceTicker]: Starting...")

	// Create WebSocket client.
	var wsDialer websocket.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil)

	if err != nil {
		println(err.Error())
	}

	// Create subscription message.
	subscribe := gdax.Message{
		Type: "subscribe",
		Channels: []gdax.MessageChannel{
			gdax.MessageChannel{
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

	message := gdax.Message{}
	for running {
		// Process messages spooled from the socket in JSON.
		if err := wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}

		if message.Type == "ticker" {
			// Write the price to channel.
			tick := PriceTick{Price: message.Price, High: message.BestAsk, Time: message.Time.Time()}
			ticker.TickChannel <- tick
		}
	}
}

func (ticker BTCPriceTicker) Stop() {
	println("[BTCPriceTicker]: Shutting down...")
	running = false
}

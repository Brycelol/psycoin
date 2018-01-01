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
	Price     float64
	High      float64
	Time time.Time
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

/*
type Ticker struct {
	TradeId int     `json:"trade_id,number"`
	Price   float64 `json:"price,string"`
	Size    float64 `json:"size,string"`
	Time    Time    `json:"time,string"`
	Bid     float64 `json:"bid,string"`
	Ask     float64 `json:"ask,string"`
	Volume  float64 `json:"volume,string"`
}
 */

 /*
 {
    "type": "ticker",
    "trade_id": 20153558,
    "sequence": 3262786978,
    "time": "2017-09-02T17:05:49.250000Z",
    "product_id": "BTC-USD",
    "price": "4388.01000000",
    "side": "buy", // Taker side
    "last_size": "0.03000000",
    "best_bid": "4388",
    "best_ask": "4388.01"
}
  */

 /*

type Message struct {
	TradeId       int              `json:"trade_id,number"`
	OrderId       string           `json:"order_id"`
	Sequence      int64            `json:"sequence,number"`
	MakerOrderId  string           `json:"maker_order_id"`
	TakerOrderId  string           `json:"taker_order_id"`
	Time          Time             `json:"time,string"`
	RemainingSize float64          `json:"remaining_size,string"`
	NewSize       float64          `json:"new_size,string"`
	OldSize       float64          `json:"old_size,string"`
	Size          float64          `json:"size,string"`
	Price         float64          `json:"price,string"`
	Side          string           `json:"side"`
	Reason        string           `json:"reason"`
	OrderType     string           `json:"order_type"`
	Funds         float64          `json:"funds,string"`
	NewFunds      float64          `json:"new_funds,string"`
	OldFunds      float64          `json:"old_funds,string"`
	Message       string           `json:"message"`
	Bids          [][]string       `json:"bids,omitempty"`
	Asks          [][]string       `json:"asks,omitempty"`
	Changes       [][]string       `json:"changes,omitempty"`
	LastSize      float64          `json:"last_size,string"`
	BestBid       float64          `json:"best_bid,string"`
	BestAsk       float64          `json:"best_ask,string"`
	Channels      []MessageChannel `json:"channels"`
}
  */

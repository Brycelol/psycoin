package trader

import ("fmt"
	"github.com/brycelol/psycoin/ticker")


type PsyTrader struct {
	TickChannel chan ticker.PriceTick
}

var traderRunning = false

func (trader *PsyTrader) Start() {
	println("[PsyTrader]: Starting...")

	traderRunning = true

	for traderRunning {
		// We have no consumers so we must read otherwise we would exhaust the channel.
		for cTick := range trader.TickChannel {
			fmt.Printf("[PsyTrader]: BTC Price [£%f] with high [£%f]  at %s\n", cTick.Price, cTick.High, cTick.Time)
		}
	}
}

func (trader *PsyTrader) Stop() {
	println("[PsyTrader]: Shutting down...")
	traderRunning = false
}

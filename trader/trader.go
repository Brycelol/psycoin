package trader

import (
	"github.com/brycelol/psycoin/ticker"
	"github.com/brycelol/psycoin/strategy"
)

func Trade(tickChannel chan ticker.PriceTick) {
	prs := strategy.PivotReversalStrategy{}

	for cTick := range tickChannel {
		prs.OnClose(cTick)
	}
}

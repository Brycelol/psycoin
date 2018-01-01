package strategy

import (
	"github.com/brycelol/psycoin/ticker"
	"fmt"
)

type Strategy interface {
	// invoked when a time period has closed.
	onClose(tick ticker.PriceTick)
}

type PivotReversalStrategy struct {
	dataCtx map[int64]float64
}

func (pvs *PivotReversalStrategy) OnClose(tick ticker.PriceTick) {
	fmt.Printf("[PivotReversalStrategy]: BTC Price [£%f] with high [£%f]  at %s\n", tick.Price, tick.High, tick.Time)

/*	if pvs.dataCtx == nil {
		pvs.dataCtx = make(map[int64]float64)
		pvs.dataCtx[tick.Time.Unix()] = tick.High
		// else if we already have an entry for this timestamp.
		// we wont!!
	} else if _, high := pvs.dataCtx[tick.Time.Unix()]; high {
		newHigh := tick.High

		if newHigh > pvs.dataCtx[tick.Time.Unix()] {
			pvs.dataCtx[tick.Time.Unix()] = newHigh
		}
	}*/
}



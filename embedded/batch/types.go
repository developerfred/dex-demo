package batch

import (
	"github.com/tendermint/dex-demo/storeutils"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/dex-demo/pkg/matcheng"
)

type Batch struct {
	BlockNumber   int64                     `json:"block_number"`
	BlockTime     time.Time                 `json:"block_time"`
	MarketID      storeutils.EntityID       `json:"market_id"`
	ClearingPrice sdk.Uint                  `json:"clearing_price"`
	Bids          []matcheng.AggregatePrice `json:"bids"`
	Asks          []matcheng.AggregatePrice `json:"asks"`
}

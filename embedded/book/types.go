package book

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/dex-demo/pkg/matcheng"
	"github.com/tendermint/dex-demo/storeutils"
)

type Book struct {
	Bids []matcheng.AggregatePrice `json:"bids"`
	Asks []matcheng.AggregatePrice `json:"asks"`
}

type QueryResultEntry struct {
	Price    sdk.Uint `json:"price"`
	Quantity sdk.Uint `json:"quantity"`
}

type QueryResult struct {
	MarketID    storeutils.EntityID `json:"market_id"`
	BlockNumber int64               `json:"block_number"`
	Bids        []QueryResultEntry  `json:"bids"`
	Asks        []QueryResultEntry  `json:"asks"`
}

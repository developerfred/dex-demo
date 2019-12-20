package types

import (
	"github.com/tendermint/dex-demo/storeutils"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/dex-demo/pkg/matcheng"
)

type EventHandler interface {
	OnEvent(event interface{}) error
}

type Batch struct {
	BlockNumber   int64
	BlockTime     time.Time
	MarketID      storeutils.EntityID
	ClearingPrice sdk.Uint
	Bids          []matcheng.AggregatePrice
	Asks          []matcheng.AggregatePrice
}

type Fill struct {
	OrderID     storeutils.EntityID
	MarketID    storeutils.EntityID
	Owner       sdk.AccAddress
	Pair        string
	Direction   matcheng.Direction
	QtyFilled   sdk.Uint
	QtyUnfilled sdk.Uint
	BlockNumber int64
	BlockTime   int64
	Price       sdk.Uint
}

type OrderCreated struct {
	ID                storeutils.EntityID
	Owner             sdk.AccAddress
	MarketID          storeutils.EntityID
	Direction         matcheng.Direction
	Price             sdk.Uint
	Quantity          sdk.Uint
	TimeInForceBlocks uint16
	CreatedBlock      int64
}

type OrderCancelled struct {
	OrderID storeutils.EntityID
}

type BurnCreated struct {
	ID          storeutils.EntityID
	AssetID     storeutils.EntityID
	BlockNumber int64
	Burner      sdk.AccAddress
	Beneficiary []byte
	Quantity    sdk.Uint
}

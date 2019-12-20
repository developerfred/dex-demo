package exchange

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/dex-demo/embedded"
	"github.com/tendermint/dex-demo/pkg/matcheng"
	"github.com/tendermint/dex-demo/storeutils"
)

type OrderCreationRequest struct {
	MarketID    storeutils.EntityID `json:"market_id"`
	Direction   matcheng.Direction  `json:"direction"`
	Price       sdk.Uint            `json:"price"`
	Quantity    sdk.Uint            `json:"quantity"`
	Type        string              `json:"type"`
	TimeInForce uint16              `json:"time_in_force"`
}

type OrderCreationResponse struct {
	BlockInclusion embedded.BlockInclusion `json:"block_inclusion"`
	ID             storeutils.EntityID     `json:"id"`
	MarketID       storeutils.EntityID     `json:"market_id"`
	Direction      matcheng.Direction      `json:"direction"`
	Price          sdk.Uint                `json:"price"`
	Quantity       sdk.Uint                `json:"quantity"`
	Type           string                  `json:"type"`
	TimeInForce    uint16                  `json:"time_in_force"`
	Status         string                  `json:"status"`
}

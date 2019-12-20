package types

import (
	"github.com/tendermint/dex-demo/storeutils"
)

const (
	ModuleName = "market"
	RouterKey  = ModuleName
	StoreKey   = RouterKey
)

type Market struct {
	ID           storeutils.EntityID
	BaseAssetID  storeutils.EntityID
	QuoteAssetID storeutils.EntityID
}

func New(id storeutils.EntityID, baseAsset storeutils.EntityID, quoteAsset storeutils.EntityID) Market {
	return Market{
		ID:           id,
		BaseAssetID:  baseAsset,
		QuoteAssetID: quoteAsset,
	}
}

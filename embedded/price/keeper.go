package price

import (
	"github.com/tendermint/dex-demo/storeutils"
	"time"

	dbm "github.com/tendermint/tm-db"

	"github.com/tendermint/dex-demo/types"
	"github.com/tendermint/dex-demo/types/store"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/store/types"
)

type IteratorCB func(tick Tick) bool

type Keeper struct {
	as  store.ArchiveStore
	cdc *codec.Codec
}

func NewKeeper(db dbm.DB, cdc *codec.Codec) Keeper {
	return Keeper{
		as:  store.NewTable(db, EntityName),
		cdc: cdc,
	}
}

func (k Keeper) ReverseIteratorByMarket(mktID storeutils.EntityID, cb IteratorCB) {
	k.as.PrefixIterator(tickIterKey(mktID), func(_ []byte, v []byte) bool {
		var tick Tick
		k.cdc.MustUnmarshalBinaryBare(v, &tick)
		return cb(tick)
	})
}

func (k Keeper) ReverseIteratorByMarketFrom(mktID storeutils.EntityID, from time.Time, cb IteratorCB) {
	k.as.ReverseIterator(tickKey(mktID, 0), sdk.PrefixEndBytes(tickKey(mktID, 0)), func(_ []byte, v []byte) bool {
		var tick Tick
		k.cdc.MustUnmarshalBinaryBare(v, &tick)
		return cb(tick)
	})
}

func (k Keeper) IteratorByMarketAndInterval(mktID storeutils.EntityID, from time.Time, to time.Time, cb IteratorCB) {
	k.as.Iterator(tickKey(mktID, from.Unix()), sdk.PrefixEndBytes(tickKey(mktID, to.Unix())), func(_ []byte, v []byte) bool {
		var tick Tick
		k.cdc.MustUnmarshalBinaryBare(v, &tick)
		return cb(tick)
	})
}

func (k Keeper) OnFillEvent(event types.Fill) {
	tick := Tick{
		MarketID:    event.MarketID,
		Pair:        event.Pair,
		BlockNumber: event.BlockNumber,
		BlockTime:   event.BlockTime,
		Price:       event.Price,
	}
	storedB := k.cdc.MustMarshalBinaryBare(tick)
	k.as.Set(tickKey(event.MarketID, tick.BlockTime), storedB)
}

func (k Keeper) OnEvent(event interface{}) error {
	switch ev := event.(type) {
	case types.Fill:
		k.OnFillEvent(ev)
	}

	return nil
}

func tickKey(mktID storeutils.EntityID, blockTime int64) []byte {
	return storeutils.PrefixKeyBytes(tickIterKey(mktID), storeutils.Int64Subkey(blockTime))
}

func tickIterKey(mktID storeutils.EntityID) []byte {
	return storeutils.PrefixKeyString("tick", mktID.Bytes())
}

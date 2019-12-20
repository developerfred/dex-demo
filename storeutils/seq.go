package storeutils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IncrementSeq(ctx sdk.Context, sk sdk.StoreKey, seqKey []byte) EntityID {
	store := ctx.KVStore(sk)
	seq := GetSeq(ctx, sk, seqKey).Inc()
	store.Set(seqKey, []byte(seq.String()))
	return seq
}

func GetSeq(ctx sdk.Context, sk sdk.StoreKey, seqKey []byte) EntityID {
	store := ctx.KVStore(sk)
	if !store.Has(seqKey) {
		return ZeroEntityID
	}

	b := store.Get(seqKey)
	return NewEntityIDFromString(string(b))
}

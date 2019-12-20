package storeutils

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

func Get(ctx sdk.Context, sk sdk.StoreKey, cdc *codec.Codec, key []byte, proto interface{}) error {
	store := ctx.KVStore(sk)
	b := store.Get(key)
	if b == nil {
		return ErrNotFound
	}
	cdc.MustUnmarshalBinaryBare(b, proto)
	return nil
}

func Set(ctx sdk.Context, sk sdk.StoreKey, cdc *codec.Codec, key []byte, val interface{}) {
	store := ctx.KVStore(sk)
	store.Set(key, cdc.MustMarshalBinaryBare(val))
}

func SetNotExists(ctx sdk.Context, sk sdk.StoreKey, cdc *codec.Codec, key []byte, val interface{}) error {
	if Has(ctx, sk, key) {
		return ErrAlreadyExists
	}
	Set(ctx, sk, cdc, key, val)
	return nil
}

func SetExists(ctx sdk.Context, sk sdk.StoreKey, cdc *codec.Codec, key []byte, val interface{}) error {
	if !Has(ctx, sk, key) {
		return ErrNotFound
	}
	Set(ctx, sk, cdc, key, val)
	return nil
}

func Del(ctx sdk.Context, sk sdk.StoreKey, key []byte) error {
	if !Has(ctx, sk, key) {
		return ErrNotFound
	}
	store := ctx.KVStore(sk)
	store.Delete(key)
	return nil
}

func Has(ctx sdk.Context, sk sdk.StoreKey, key []byte) bool {
	store := ctx.KVStore(sk)
	return store.Has(key)
}

package store

import (
	"fmt"
	"github.com/tendermint/dex-demo/storeutils"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func FormatCoin(id storeutils.EntityID, amount sdk.Uint) sdk.Coin {
	out, err := sdk.ParseCoin(fmt.Sprintf("%s%s", amount.String(), FormatDenom(id)))
	// should never happen
	if err != nil {
		panic(err)
	}
	return out
}

func FormatDenom(id storeutils.EntityID) string {
	return fmt.Sprintf("asset%s", id.String())
}

package matcheng

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/dex-demo/storeutils"
)

type Fill struct {
	OrderID     storeutils.EntityID
	QtyFilled   sdk.Uint
	QtyUnfilled sdk.Uint
}

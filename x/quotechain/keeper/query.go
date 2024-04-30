package keeper

import (
	"quotechain/x/quotechain/types"
)

var _ types.QueryServer = Keeper{}

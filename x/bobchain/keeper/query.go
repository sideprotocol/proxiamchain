package keeper

import (
	"bobchain/x/bobchain/types"
)

var _ types.QueryServer = Keeper{}

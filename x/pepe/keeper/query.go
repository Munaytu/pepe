package keeper

import (
	"pepe/x/pepe/types"
)

var _ types.QueryServer = Keeper{}

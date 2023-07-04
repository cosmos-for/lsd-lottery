package keeper

import (
	"lsd-lottery/x/wasm/types"
)

var _ types.QueryServer = Keeper{}

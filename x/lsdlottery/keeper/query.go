package keeper

import (
	"lsd-lottery/x/lsdlottery/types"
)

var _ types.QueryServer = Keeper{}

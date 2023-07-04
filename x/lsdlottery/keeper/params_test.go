package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "lsd-lottery/testutil/keeper"
	"lsd-lottery/x/lsdlottery/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LsdlotteryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

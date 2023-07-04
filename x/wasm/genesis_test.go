package wasm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lsd-lottery/testutil/keeper"
	"lsd-lottery/testutil/nullify"
	"lsd-lottery/x/wasm"
	"lsd-lottery/x/wasm/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WasmKeeper(t)
	wasm.InitGenesis(ctx, *k, genesisState)
	got := wasm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

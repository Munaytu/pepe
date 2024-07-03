package pepe_test

import (
	"testing"

	keepertest "pepe/testutil/keeper"
	"pepe/testutil/nullify"
	pepe "pepe/x/pepe/module"
	"pepe/x/pepe/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PepeKeeper(t)
	pepe.InitGenesis(ctx, k, genesisState)
	got := pepe.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

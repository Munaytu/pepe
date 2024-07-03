package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "pepe/testutil/keeper"
	"pepe/x/pepe/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PepeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

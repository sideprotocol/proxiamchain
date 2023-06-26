package bobchain_test

import (
	"testing"

	keepertest "bobchain/testutil/keeper"
	"bobchain/testutil/nullify"
	"bobchain/x/bobchain"
	"bobchain/x/bobchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BobchainKeeper(t)
	bobchain.InitGenesis(ctx, *k, genesisState)
	got := bobchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

package quotechain_test

import (
	"testing"

	keepertest "quotechain/testutil/keeper"
	"quotechain/testutil/nullify"
	quotechain "quotechain/x/quotechain/module"
	"quotechain/x/quotechain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.QuotechainKeeper(t)
	quotechain.InitGenesis(ctx, k, genesisState)
	got := quotechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

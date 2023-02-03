package stellartalk_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "stellartalk/testutil/keeper"
	"stellartalk/testutil/nullify"
	"stellartalk/x/stellartalk"
	"stellartalk/x/stellartalk/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ChatList: []types.Chat{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ChatCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StellartalkKeeper(t)
	stellartalk.InitGenesis(ctx, *k, genesisState)
	got := stellartalk.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ChatList, got.ChatList)
	require.Equal(t, genesisState.ChatCount, got.ChatCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

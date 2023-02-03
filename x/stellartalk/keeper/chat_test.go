package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "stellartalk/testutil/keeper"
	"stellartalk/testutil/nullify"
	"stellartalk/x/stellartalk/keeper"
	"stellartalk/x/stellartalk/types"
)

func createNChat(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Chat {
	items := make([]types.Chat, n)
	for i := range items {
		items[i].Id = keeper.AppendChat(ctx, items[i])
	}
	return items
}

func TestChatGet(t *testing.T) {
	keeper, ctx := keepertest.StellartalkKeeper(t)
	items := createNChat(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetChat(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestChatRemove(t *testing.T) {
	keeper, ctx := keepertest.StellartalkKeeper(t)
	items := createNChat(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChat(ctx, item.Id)
		_, found := keeper.GetChat(ctx, item.Id)
		require.False(t, found)
	}
}

func TestChatGetAll(t *testing.T) {
	keeper, ctx := keepertest.StellartalkKeeper(t)
	items := createNChat(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllChat(ctx)),
	)
}

func TestChatCount(t *testing.T) {
	keeper, ctx := keepertest.StellartalkKeeper(t)
	items := createNChat(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetChatCount(ctx))
}

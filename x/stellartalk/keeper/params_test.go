package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "stellartalk/testutil/keeper"
	"stellartalk/x/stellartalk/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.StellartalkKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

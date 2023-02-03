package stellartalk

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"stellartalk/x/stellartalk/keeper"
	"stellartalk/x/stellartalk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the chat
	for _, elem := range genState.ChatList {
		k.SetChat(ctx, elem)
	}

	// Set chat count
	k.SetChatCount(ctx, genState.ChatCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ChatList = k.GetAllChat(ctx)
	genesis.ChatCount = k.GetChatCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

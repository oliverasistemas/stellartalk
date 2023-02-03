package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"stellartalk/x/stellartalk/types"
)

func (k msgServer) CreateChat(goCtx context.Context, msg *types.MsgCreateChat) (*types.MsgCreateChatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var chat = types.Chat{
		Content:   msg.Content,
		Recipient: msg.Recipient,
		Sender:    msg.Creator,
	}

	_, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		panic(err)
	}

	k.AppendChat(ctx, chat)

	return &types.MsgCreateChatResponse{}, nil
}

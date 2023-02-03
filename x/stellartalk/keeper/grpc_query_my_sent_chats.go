package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stellartalk/x/stellartalk/types"
)

func (k Keeper) MySentChats(goCtx context.Context, req *types.QueryMySentChatsRequest) (*types.QueryMySentChatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var chats []types.Chat

	store := ctx.KVStore(k.storeKey)
	chatStore := prefix.NewStore(store, types.KeyPrefix(types.ChatKey))

	pageRes, err := query.Paginate(chatStore, req.Pagination, func(key []byte, value []byte) error {
		var chat types.Chat
		if err := k.cdc.Unmarshal(value, &chat); err != nil {
			return err
		}

		chats = append(chats, chat)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryMySentChatsResponse{Chat: chats, Pagination: pageRes}, nil

}

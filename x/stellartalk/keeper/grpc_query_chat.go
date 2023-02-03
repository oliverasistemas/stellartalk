package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stellartalk/x/stellartalk/types"
)

func (k Keeper) ChatAll(c context.Context, req *types.QueryAllChatRequest) (*types.QueryAllChatResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var chats []types.Chat
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	chatStore := prefix.NewStore(store, types.KeyPrefix(types.ChatKey))

	pageRes, err := query.Paginate(chatStore, req.Pagination, func(key []byte, value []byte) error {
		var chat types.Chat
		if err := k.cdc.Unmarshal(value, &chat); err != nil {
			return err
		}

		//Filter sender
		if req.Sender != "" && chat.Sender != req.Sender {
			return nil
		}

		//Filter recipient
		if req.Recipient != "" && chat.Recipient != req.Recipient {
			return nil
		}

		chats = append(chats, chat)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChatResponse{Chat: chats, Pagination: pageRes}, nil
}

func (k Keeper) Chat(c context.Context, req *types.QueryGetChatRequest) (*types.QueryGetChatResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	chat, found := k.GetChat(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetChatResponse{Chat: chat}, nil
}

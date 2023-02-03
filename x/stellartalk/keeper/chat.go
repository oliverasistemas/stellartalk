package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"stellartalk/x/stellartalk/types"
)

// GetChatCount get the total number of chat
func (k Keeper) GetChatCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ChatCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetChatCount set the total number of chat
func (k Keeper) SetChatCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ChatCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendChat appends a chat in the store with a new id and update the count
func (k Keeper) AppendChat(
	ctx sdk.Context,
	chat types.Chat,
) uint64 {
	// Create the chat
	count := k.GetChatCount(ctx)

	// Set the ID of the appended value
	chat.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChatKey))
	appendedValue := k.cdc.MustMarshal(&chat)
	store.Set(GetChatIDBytes(chat.Id), appendedValue)

	// Update chat count
	k.SetChatCount(ctx, count+1)

	return count
}

// SetChat set a specific chat in the store
func (k Keeper) SetChat(ctx sdk.Context, chat types.Chat) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChatKey))
	b := k.cdc.MustMarshal(&chat)
	store.Set(GetChatIDBytes(chat.Id), b)
}

// GetChat returns a chat from its id
func (k Keeper) GetChat(ctx sdk.Context, id uint64) (val types.Chat, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChatKey))
	b := store.Get(GetChatIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChat removes a chat from the store
func (k Keeper) RemoveChat(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChatKey))
	store.Delete(GetChatIDBytes(id))
}

// GetAllChat returns all chat
func (k Keeper) GetAllChat(ctx sdk.Context) (list []types.Chat) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChatKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Chat
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetChatIDBytes returns the byte representation of the ID
func GetChatIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetChatIDFromBytes returns ID in uint64 format from a byte array
func GetChatIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

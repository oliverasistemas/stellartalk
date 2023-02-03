package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateChat = "create_chat"

var _ sdk.Msg = &MsgCreateChat{}

func NewMsgCreateChat(creator string, content string, recipient string) *MsgCreateChat {
	return &MsgCreateChat{
		Creator:   creator,
		Content:   content,
		Recipient: recipient,
	}
}

func (msg *MsgCreateChat) Route() string {
	return RouterKey
}

func (msg *MsgCreateChat) Type() string {
	return TypeMsgCreateChat
}

func (msg *MsgCreateChat) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChat) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChat) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgIssueDenom = "issue_denom"

var _ sdk.Msg = &MsgIssueDenom{}

func NewMsgIssueDenom(id string, name string, schema string, sender string) *MsgIssueDenom {
	return &MsgIssueDenom{
		Id:     id,
		Name:   name,
		Schema: schema,
		Sender: sender,
	}
}

func (msg *MsgIssueDenom) Route() string {
	return RouterKey
}

func (msg *MsgIssueDenom) Type() string {
	return TypeMsgIssueDenom
}

func (msg *MsgIssueDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIssueDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssueDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

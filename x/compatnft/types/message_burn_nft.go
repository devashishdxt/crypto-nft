package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnNFT = "burn_nft"

var _ sdk.Msg = &MsgBurnNFT{}

func NewMsgBurnNFT(id string, denomId string, sender string) *MsgBurnNFT {
	return &MsgBurnNFT{
		Id:      id,
		DenomId: denomId,
		Sender:  sender,
	}
}

func (msg *MsgBurnNFT) Route() string {
	return RouterKey
}

func (msg *MsgBurnNFT) Type() string {
	return TypeMsgBurnNFT
}

func (msg *MsgBurnNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferNFT = "transfer_nft"

var _ sdk.Msg = &MsgTransferNFT{}

func NewMsgTransferNFT(id string, denomId string, sender string, recipient string) *MsgTransferNFT {
	return &MsgTransferNFT{
		Id:        id,
		DenomId:   denomId,
		Sender:    sender,
		Recipient: recipient,
	}
}

func (msg *MsgTransferNFT) Route() string {
	return RouterKey
}

func (msg *MsgTransferNFT) Type() string {
	return TypeMsgTransferNFT
}

func (msg *MsgTransferNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

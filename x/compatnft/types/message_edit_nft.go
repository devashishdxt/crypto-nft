package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEditNFT = "edit_nft"

var _ sdk.Msg = &MsgEditNFT{}

func NewMsgEditNFT(id string, denomId string, name string, uri string, data string, sender string) *MsgEditNFT {
	return &MsgEditNFT{
		Id:      id,
		DenomId: denomId,
		Name:    name,
		URI:     uri,
		Data:    data,
		Sender:  sender,
	}
}

func (msg *MsgEditNFT) Route() string {
	return RouterKey
}

func (msg *MsgEditNFT) Type() string {
	return TypeMsgEditNFT
}

func (msg *MsgEditNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEditNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

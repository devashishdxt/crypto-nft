package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintNFT = "mint_nft"

var _ sdk.Msg = &MsgMintNFT{}

func NewMsgMintNft(id string, denomId string, name string, uri string, data string, sender string, recipient string) *MsgMintNFT {
	return &MsgMintNFT{
		Id:        id,
		DenomId:   denomId,
		Name:      name,
		URI:       uri,
		Data:      data,
		Sender:    sender,
		Recipient: recipient,
	}
}

func (msg *MsgMintNFT) Route() string {
	return RouterKey
}

func (msg *MsgMintNFT) Type() string {
	return TypeMsgMintNFT
}

func (msg *MsgMintNFT) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

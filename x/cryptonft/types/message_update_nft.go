package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateNFT = "update_nft"

var _ sdk.Msg = &MsgUpdateNFT{}

func NewMsgUpdateNFT(creator string, classId string, nftId string, uri string, uriHash string, data *codectypes.Any) *MsgUpdateNFT {
	return &MsgUpdateNFT{
		Creator: creator,
		ClassId: classId,
		NftId:   nftId,
		Uri:     uri,
		UriHash: uriHash,
		Data:    data,
	}
}

func (msg *MsgUpdateNFT) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNFT) Type() string {
	return TypeMsgUpdateNFT
}

func (msg *MsgUpdateNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

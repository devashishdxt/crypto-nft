package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewClass = "new_class"

var _ sdk.Msg = &MsgNewClass{}

func NewMsgNewClass(creator string, id string, name string, symbol string, description string, uri string, uriHash string, mintRestricted bool, burnRestricted bool, updateRestricted bool, data *codectypes.Any) *MsgNewClass {
	return &MsgNewClass{
		Creator:          creator,
		Id:               id,
		Name:             name,
		Symbol:           symbol,
		Description:      description,
		Uri:              uri,
		UriHash:          uriHash,
		MintRestricted:   mintRestricted,
		BurnRestricted:   burnRestricted,
		UpdateRestricted: updateRestricted,
		Data:             data,
	}
}

func (msg *MsgNewClass) Route() string {
	return RouterKey
}

func (msg *MsgNewClass) Type() string {
	return TypeMsgNewClass
}

func (msg *MsgNewClass) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

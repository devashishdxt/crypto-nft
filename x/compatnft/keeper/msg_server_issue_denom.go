package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
)

func (k msgServer) IssueDenom(goCtx context.Context, msg *types.MsgIssueDenom) (*types.MsgIssueDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := k.AddDenom(ctx, sender, msg.Id, msg.Name, msg.Schema); err != nil {
		return nil, err
	}

	return &types.MsgIssueDenomResponse{}, nil
}

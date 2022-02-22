package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
)

func (k msgServer) EditNFT(goCtx context.Context, msg *types.MsgEditNFT) (*types.MsgEditNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := k.Edit(ctx, sender, msg.Id, msg.DenomId, msg.Name, msg.URI, msg.Data); err != nil {
		return nil, err
	}

	return &types.MsgEditNFTResponse{}, nil
}

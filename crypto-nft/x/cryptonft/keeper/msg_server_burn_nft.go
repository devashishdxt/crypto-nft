package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

func (k msgServer) BurnNFT(goCtx context.Context, msg *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := k.Burn(ctx, creator, msg.ClassId, msg.NftId); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&nft.EventBurn{
		ClassId: msg.ClassId,
		Id:      msg.NftId,
		Owner:   msg.Creator,
	})

	return &types.MsgBurnNFTResponse{}, nil
}

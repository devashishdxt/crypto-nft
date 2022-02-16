package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

func (k msgServer) MintNFT(goCtx context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	if err := k.Mint(ctx, creator, msg.ClassId, msg.Id, msg.Uri, msg.UriHash, receiver, msg.Data); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&nft.EventMint{
		ClassId: msg.ClassId,
		Id:      msg.Id,
		Owner:   msg.Receiver,
	})

	return &types.MsgMintNFTResponse{}, nil
}

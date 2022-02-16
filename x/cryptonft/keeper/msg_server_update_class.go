package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

func (k msgServer) UpdateClass(goCtx context.Context, msg *types.MsgUpdateClass) (*types.MsgUpdateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := k.PutClass(ctx, creator, msg.Id, msg.Name, msg.Symbol, msg.Description, msg.Uri, msg.UriHash, msg.MintRestricted, msg.BurnRestricted, msg.UpdateRestricted, msg.Data); err != nil {
		return nil, err
	}

	return &types.MsgUpdateClassResponse{}, nil
}

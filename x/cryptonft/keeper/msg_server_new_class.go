package keeper

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

func (k msgServer) NewClass(goCtx context.Context, msg *types.MsgNewClass) (*types.MsgNewClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	metadata := &types.ClassMetadata{
		Creator: msg.Creator,
		Data:    msg.Data,
	}

	data, err := codectypes.NewAnyWithValue(metadata)
	if err != nil {
		return nil, err
	}

	class := nft.Class{
		Id:          msg.Id,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Description: msg.Description,
		Uri:         msg.Uri,
		UriHash:     msg.UriHash,
		Data:        data,
	}

	k.nftKeeper.SaveClass(ctx, class)

	return &types.MsgNewClassResponse{}, nil
}

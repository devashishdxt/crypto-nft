package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NFT(goCtx context.Context, req *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	nft, err := k.GetNFT(ctx, req.DenomId, req.TokenId)
	if err != nil {
		return nil, err
	}

	return &types.QueryNFTResponse{
		NFT: nft,
	}, nil
}

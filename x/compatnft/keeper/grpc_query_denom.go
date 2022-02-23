package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Denom(goCtx context.Context, req *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	denom, err := k.GetDenom(ctx, req.DenomId)
	if err != nil {
		return nil, err
	}

	return &types.QueryDenomResponse{
		Denom: denom,
	}, nil
}

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Class(goCtx context.Context, req *types.QueryClassRequest) (*types.QueryClassResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	class, found := k.GetClass(ctx, req.ClassId)
	if !found {
		return nil, status.Error(codes.NotFound, "class not found")
	}

	return &types.QueryClassResponse{
		Class: class,
	}, nil
}

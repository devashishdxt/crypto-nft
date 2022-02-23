package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Classes(goCtx context.Context, req *types.QueryClassesRequest) (*types.QueryClassesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	classes, err := k.NftKeeper().Classes(goCtx, &nft.QueryClassesRequest{
		Pagination: req.Pagination,
	})
	if err != nil {
		return nil, err
	}

	var resultClasses []*types.Class
	for _, class := range classes.Classes {
		resultClass := k.ConvertClass(class)
		resultClasses = append(resultClasses, resultClass)
	}

	return &types.QueryClassesResponse{
		Classes:    resultClasses,
		Pagination: classes.Pagination,
	}, nil
}

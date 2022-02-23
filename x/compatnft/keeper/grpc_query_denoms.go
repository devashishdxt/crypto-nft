package keeper

import (
	"context"

	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	cryptonfttypes "github.com/devashishdxt/crypto-nft/x/cryptonft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Denoms(goCtx context.Context, req *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	classes, err := k.CryptoNftKeeper().Classes(goCtx, &cryptonfttypes.QueryClassesRequest{
		Pagination: req.Pagination,
	})
	if err != nil {
		return nil, err
	}

	var denoms []types.Denom
	for _, class := range classes.Classes {
		denom, err := k.ConvertClass(class)
		if err != nil {
			return nil, err
		}
		denoms = append(denoms, *denom)
	}

	return &types.QueryDenomsResponse{
		Denoms:     denoms,
		Pagination: classes.Pagination,
	}, nil
}

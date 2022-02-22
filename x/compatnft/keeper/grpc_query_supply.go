package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Supply(goCtx context.Context, req *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if len(req.Owner) == 0 && len(req.DenomId) > 0 {
		supplyResponse, err := k.NftKeeper().Supply(goCtx, &nft.QuerySupplyRequest{
			ClassId: req.DenomId,
		})
		if err != nil {
			return nil, err
		}

		return &types.QuerySupplyResponse{
			Amount: supplyResponse.Amount,
		}, nil
	} else {
		balanceResposne, err := k.NftKeeper().Balance(goCtx, &nft.QueryBalanceRequest{
			ClassId: req.DenomId,
			Owner:   req.Owner,
		})

		if err != nil {
			return nil, err
		}

		return &types.QuerySupplyResponse{
			Amount: balanceResposne.Amount,
		}, nil
	}
}

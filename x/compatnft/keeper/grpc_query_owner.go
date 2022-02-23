package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Owner(goCtx context.Context, req *types.QueryOwnerRequest) (*types.QueryOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	nftsOfClass, err := k.NftKeeper().NFTs(goCtx, &nft.QueryNFTsRequest{
		ClassId:    req.DenomId,
		Owner:      req.Owner,
		Pagination: req.Pagination,
	})

	if err != nil {
		return nil, err
	}

	var denomMap = make(map[string][]string)
	var denoms = make([]string, 0, len(nftsOfClass.Nfts))
	for _, token := range nftsOfClass.Nfts {
		denomMap[token.ClassId] = append(denomMap[token.ClassId], token.Id)
		denoms = append(denoms, token.ClassId)
	}

	var idc []types.IDCollection
	for _, denomId := range denoms {
		idc = append(idc, types.IDCollection{DenomId: denomId, TokenIds: denomMap[denomId]})
	}

	return &types.QueryOwnerResponse{
		Owner: &types.Owner{
			Address:       req.Owner,
			IDCollections: idc,
		},
		Pagination: nftsOfClass.Pagination,
	}, nil
}

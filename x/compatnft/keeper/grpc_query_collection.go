package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Collection(goCtx context.Context, req *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	class, found := k.CryptoNftKeeper().GetClass(ctx, req.DenomId)
	if !found {
		return nil, status.Error(codes.NotFound, "denom not found")
	}

	nftsOfClass, err := k.NftKeeper().NFTs(goCtx, &nft.QueryNFTsRequest{
		ClassId:    req.DenomId,
		Pagination: req.Pagination,
	})
	if err != nil {
		return nil, err
	}

	var nfts []types.BaseNFT
	for _, token := range nftsOfClass.Nfts {
		owner := k.NftKeeper().GetOwner(ctx, req.DenomId, token.Id)

		var nftMetadata types.NFTMetadata
		if err := k.cdc.Unmarshal(token.Data.GetValue(), &nftMetadata); err != nil {
			return nil, err
		}

		nfts = append(nfts, types.BaseNFT{
			Id:    token.Id,
			URI:   token.Uri,
			Name:  nftMetadata.Name,
			Owner: owner.String(),
			Data:  nftMetadata.Data,
		})
	}

	var denomMetadata types.DenomMetadata
	if err := k.cdc.Unmarshal(class.Data.GetValue(), &denomMetadata); err != nil {
		return nil, err
	}

	collection := &types.Collection{
		Denom: types.Denom{
			Id:      class.Id,
			Name:    class.Name,
			Schema:  denomMetadata.Schema,
			Creator: class.Creator,
		},
		NFTs: nfts,
	}

	return &types.QueryCollectionResponse{
		Collection: collection,
		Pagination: nftsOfClass.Pagination,
	}, nil
}

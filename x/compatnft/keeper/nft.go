package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
)

func (k Keeper) Mint(ctx sdk.Context, sender sdk.AccAddress, nftId string, denomId string, name string, uri string, data string, receiver sdk.AccAddress) error {
	nftMetadata := &types.NFTMetadata{
		Name: name,
		Data: data,
	}

	metadata, err := codectypes.NewAnyWithValue(nftMetadata)
	if err != nil {
		return err
	}

	return k.CryptoNftKeeper().Mint(ctx, sender, denomId, nftId, uri, "", receiver, metadata)
}

func (k Keeper) Edit(ctx sdk.Context, sender sdk.AccAddress, nftId string, denomId string, name string, uri string, data string) error {
	nftMetadata := &types.NFTMetadata{
		Data: data,
	}

	metadata, err := codectypes.NewAnyWithValue(nftMetadata)
	if err != nil {
		return err
	}

	return k.CryptoNftKeeper().Update(ctx, sender, denomId, nftId, uri, "", metadata)
}

func (k Keeper) Transfer(ctx sdk.Context, sender sdk.AccAddress, nftId string, denomId string, recipient sdk.AccAddress) error {
	return k.CryptoNftKeeper().Transfer(ctx, sender, denomId, nftId, recipient)
}

func (k Keeper) Burn(ctx sdk.Context, sender sdk.AccAddress, nftId string, denomId string) error {
	return k.CryptoNftKeeper().Burn(ctx, sender, denomId, nftId)
}

func (k Keeper) GetNFT(ctx sdk.Context, denomId string, tokenId string) (*types.BaseNFT, error) {
	nft, found := k.NftKeeper().GetNFT(ctx, denomId, tokenId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNftNotFound, "denomId: %s, tokenId: %s", denomId, tokenId)
	}

	owner := k.NftKeeper().GetOwner(ctx, denomId, tokenId)

	var nftMetadata types.NFTMetadata
	if err := k.cdc.Unmarshal(nft.Data.GetValue(), &nftMetadata); err != nil {
		return nil, err
	}

	return &types.BaseNFT{
		Id:    nft.Id,
		Name:  nftMetadata.Name,
		URI:   nft.Uri,
		Data:  nftMetadata.Data,
		Owner: owner.String(),
	}, nil
}

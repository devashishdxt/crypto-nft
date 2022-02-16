package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

// Mint defines a method for minting a new nft
func (k Keeper) Mint(ctx sdk.Context, creator sdk.AccAddress, classId string, id string, uri string, uriHash string, receiver sdk.AccAddress, data *codectypes.Any) error {
	if err := k.CanMint(ctx, creator, classId); err != nil {
		return err
	}

	nft := nft.NFT{
		ClassId: classId,
		Id:      id,
		Uri:     uri,
		UriHash: uriHash,
		Data:    data,
	}

	return k.nftKeeper.Mint(ctx, nft, receiver)
}

// CanMint return error when the creator can't mint a new nft of given class
func (k Keeper) CanMint(ctx sdk.Context, creator sdk.AccAddress, classId string) error {
	class, found := k.nftKeeper.GetClass(ctx, classId)
	if !found {
		return sdkerrors.Wrap(nft.ErrClassNotExists, classId)
	}

	var classMetadata types.ClassMetadata
	if err := k.cdc.Unmarshal(class.Data.GetValue(), &classMetadata); err != nil {
		return err
	}

	if classMetadata.MintRestricted && classMetadata.Creator != creator.String() {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to mint NFT of class %s", creator, classId)
	}

	return nil
}

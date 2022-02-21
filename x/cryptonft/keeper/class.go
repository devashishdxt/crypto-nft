package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

// AddClass defines a method for creating a new NFT classification
func (k Keeper) AddClass(ctx sdk.Context, creator sdk.AccAddress, id string, name string, symbol string, description string, uri string, uriHash string, mintRestricted bool, burnRestricted bool, updateRestricted bool, data *codectypes.Any) error {
	metadata := &types.ClassMetadata{
		Creator:          creator.String(),
		MintRestricted:   mintRestricted,
		BurnRestricted:   burnRestricted,
		UpdateRestricted: updateRestricted,
		Data:             data,
	}

	data, err := codectypes.NewAnyWithValue(metadata)
	if err != nil {
		return err
	}

	class := nft.Class{
		Id:          id,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         uri,
		UriHash:     uriHash,
		Data:        data,
	}

	k.nftKeeper.SaveClass(ctx, class)

	return nil
}

// PutClass defines a method for updating a exist nft class
func (k Keeper) PutClass(ctx sdk.Context, creator sdk.AccAddress, id string, name string, symbol string, description string, uri string, uriHash string, mintRestricted bool, burnRestricted bool, updateRestricted bool, data *codectypes.Any) error {
	class, found := k.nftKeeper.GetClass(ctx, id)
	if !found {
		return sdkerrors.Wrap(nft.ErrClassNotExists, id)
	}

	var classMetadata types.ClassMetadata
	if err := k.cdc.Unmarshal(class.Data.GetValue(), &classMetadata); err != nil {
		return err
	}

	if classMetadata.Creator != creator.String() {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to update class %s", creator, id)
	}

	metadata := &types.ClassMetadata{
		Creator:          creator.String(),
		MintRestricted:   mintRestricted,
		BurnRestricted:   burnRestricted,
		UpdateRestricted: updateRestricted,
		Data:             data,
	}

	data, err := codectypes.NewAnyWithValue(metadata)
	if err != nil {
		return err
	}

	class = nft.Class{
		Id:          id,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         uri,
		UriHash:     uriHash,
		Data:        data,
	}

	k.nftKeeper.UpdateClass(ctx, class)

	return nil
}

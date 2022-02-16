package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

// SaveClass defines a method for creating a new NFT classification
func (k Keeper) SaveClass(ctx sdk.Context, creator sdk.AccAddress, id string, name string, symbol string, description string, uri string, uriHash string, mintRestricted bool, burnRestricted bool, updateRestricted bool, data *codectypes.Any) error {
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

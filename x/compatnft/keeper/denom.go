package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
)

func (k Keeper) AddDenom(ctx sdk.Context, sender sdk.AccAddress, id string, name string, schema string) error {
	denomMetadata := &types.DenomMetadata{
		Schema: schema,
	}

	data, err := codectypes.NewAnyWithValue(denomMetadata)
	if err != nil {
		return err
	}

	return k.CryptoNftKeeper().AddClass(ctx, sender, id, name, "", "", "", "", true, true, true, data)
}

func (k Keeper) GetDenom(ctx sdk.Context, denomId string) (*types.Denom, error) {
	class, found := k.CryptoNftKeeper().GetClass(ctx, denomId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDenomNotFound, "not found denomId: %s", denomId)
	}

	var denomMetadata types.DenomMetadata
	if err := k.cdc.Unmarshal(class.Data.GetValue(), &denomMetadata); err != nil {
		return nil, err
	}

	return &types.Denom{
		Id:      class.Id,
		Name:    class.Name,
		Schema:  denomMetadata.Schema,
		Creator: class.Creator,
	}, nil
}

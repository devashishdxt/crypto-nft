package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

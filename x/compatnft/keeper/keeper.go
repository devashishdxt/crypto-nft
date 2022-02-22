package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	cryptonftkeeper "github.com/devashishdxt/crypto-nft/x/cryptonft/keeper"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		cryptoNftKeeper cryptonftkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	cryptoNftKeeper cryptonftkeeper.Keeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		cryptoNftKeeper: cryptoNftKeeper,
	}
}

func (k Keeper) CryptoNftKeeper() cryptonftkeeper.Keeper {
	return k.cryptoNftKeeper
}

func (k Keeper) NftKeeper() nftkeeper.Keeper {
	return k.CryptoNftKeeper().NftKeeper()
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

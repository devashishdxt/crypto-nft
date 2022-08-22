package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/devashishdxt/crypto-nft/x/compatnft/keeper"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"github.com/stretchr/testify/require"
	tmdb "github.com/tendermint/tm-db"
)

func CompatnftKeeper(t testing.TB, accountKeeper nft.AccountKeeper, bankKeeper nft.BankKeeper) (*keeper.Keeper, sdk.Context) {
	cryptoNftKeeper, ctx := CryptonftKeeper(t, accountKeeper, bankKeeper)

	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := ctx.MultiStore().(sdk.CommitMultiStore)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CompatnftParams",
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		*cryptoNftKeeper,
	)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/devashishdxt/crypto-nft/x/cryptonft/types"
    "github.com/devashishdxt/crypto-nft/x/cryptonft/keeper"
    keepertest "github.com/devashishdxt/crypto-nft/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CryptonftKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

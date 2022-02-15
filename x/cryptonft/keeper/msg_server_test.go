package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/devashishdxt/crypto-nft/testutil/keeper"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/keeper"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CryptonftKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

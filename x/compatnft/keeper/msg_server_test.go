package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/devashishdxt/crypto-nft/testutil/keeper"
	"github.com/devashishdxt/crypto-nft/x/compatnft/keeper"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CompatnftKeeper(t, nil, nil)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

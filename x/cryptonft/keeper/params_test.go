package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/devashishdxt/crypto-nft/testutil/keeper"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CryptonftKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

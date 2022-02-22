package keeper_test

import (
	"testing"

	testkeeper "github.com/devashishdxt/crypto-nft/testutil/keeper"
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CompatnftKeeper(t, nil, nil)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

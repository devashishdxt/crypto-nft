package cryptonft_test

import (
	"testing"

	keepertest "github.com/devashishdxt/crypto-nft/testutil/keeper"
	"github.com/devashishdxt/crypto-nft/testutil/nullify"
	"github.com/devashishdxt/crypto-nft/x/cryptonft"
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CryptonftKeeper(t)
	cryptonft.InitGenesis(ctx, *k, genesisState)
	got := cryptonft.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}

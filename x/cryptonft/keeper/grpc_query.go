package keeper

import (
	"github.com/devashishdxt/crypto-nft/x/cryptonft/types"
)

var _ types.QueryServer = Keeper{}

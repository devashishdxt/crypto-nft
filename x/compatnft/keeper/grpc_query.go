package keeper

import (
	"github.com/devashishdxt/crypto-nft/x/compatnft/types"
)

var _ types.QueryServer = Keeper{}

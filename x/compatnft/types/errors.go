package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/compatnft module sentinel errors
var (
	ErrDenomNotFound = sdkerrors.Register(ModuleName, 1, "denom not found")
)

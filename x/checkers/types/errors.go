package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1100, "invalid black address")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1101, "invalid red address")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1102, "game not parseable")
)

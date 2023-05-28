package types

import (
	"fmt"

	"github.com/alice/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s StoredGame) GetBlackAddress() (sdk.AccAddress, error) {
	black, err := sdk.AccAddressFromBech32(s.Black)
	return black, sdkerrors.Wrapf(err, ErrInvalidBlack.Error())
}

func (s StoredGame) ParseGame() (*rules.Game, error) {
	board, err := rules.Parse(s.Board)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[s.Turn].Player
	if board.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(fmt.Errorf("turn: %s", s.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

func (s StoredGame) GetRedAddress() (sdk.AccAddress, error) {
	red, err := sdk.AccAddressFromBech32(s.Red)
	return red, sdkerrors.Wrapf(err, ErrInvalidRed.Error())
}

func (s StoredGame) Validate() error {
	_, err := s.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = s.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = s.ParseGame()
	return err
}

package keeper

import (
	"context"
	"strconv"

	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// fetch game info from store
	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
	}

	// check that the game has not finished yet
	if storedGame.Winner != rules.PieceStrings[rules.NO_PLAYER] {
		return nil, types.ErrGameFinished
	}

	// check if player is legitimate
	isBlack := storedGame.Black == msg.Creator
	isRed := storedGame.Red == msg.Creator
	var player rules.Player
	if !isBlack && !isRed {
		return nil, sdkerrors.Wrapf(types.ErrCreatorNotPlayer, "%s", msg.Creator)
	} else if isBlack && isRed {
		player = rules.StringPieces[storedGame.Turn].Player
	} else if isBlack {
		player = rules.BLACK_PLAYER
	} else {
		player = rules.RED_PLAYER
	}

	// instantiate board to implement the rules
	game, err := storedGame.ParseGame()
	if err != nil {
		panic(err.Error())
	}

	// check if it is the player's turn
	if !game.TurnIs(player) {
		return nil, sdkerrors.Wrapf(types.ErrNotPlayerTurn, "%s", player)
	}

	// collect wager when player plays for the first time
	err = k.Keeper.CollectWager(ctx, &storedGame)
	if err != nil {
		return nil, err
	}

	// actually carry out the move
	captured, moveErr := game.Move(
		rules.Pos{
			X: int(msg.FromX),
			Y: int(msg.FromY),
		},
		rules.Pos{
			X: int(msg.ToX),
			Y: int(msg.ToY),
		},
	)
	if moveErr != nil {
		return nil, sdkerrors.Wrapf(types.ErrWrongMove, moveErr.Error())
	}

	// update the board

	storedGame.Winner = rules.PieceStrings[game.Winner()]

	lastBoard := game.String()

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	if storedGame.Winner == rules.PieceStrings[rules.NO_PLAYER] {
		storedGame.Board = lastBoard
		k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)
	} else {
		storedGame.Board = ""
		k.Keeper.RemoveFromFifo(ctx, &storedGame, &systemInfo)
		k.Keeper.MustPayWinnings(ctx, &storedGame)

		// Here you can register a win
		k.Keeper.MustRegisterPlayerWin(ctx, &storedGame)
	}

	storedGame.Deadline = types.FormatDeadline(types.GetNextDeadline(ctx))
	storedGame.MoveCount++
	storedGame.Turn = rules.PieceStrings[game.Turn]
	k.Keeper.SetStoredGame(ctx, storedGame)
	k.Keeper.SetSystemInfo(ctx, systemInfo)
	ctx.GasMeter().ConsumeGas(types.PlayMoveGas, "Play a move")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.MovePlayedEventType,
			sdk.NewAttribute(types.MovePlayedEventCreator, msg.Creator),
			sdk.NewAttribute(types.MovePlayedEventGameIndex, msg.GameIndex),
			sdk.NewAttribute(types.MovePlayedEventCapturedX, strconv.FormatInt(int64(captured.X), 10)),
			sdk.NewAttribute(types.MovePlayedEventCapturedY, strconv.FormatInt(int64(captured.Y), 10)),
			sdk.NewAttribute(types.MovePlayedEventWinner, rules.PieceStrings[game.Winner()]),
			sdk.NewAttribute(types.MovePlayedEventBoard, lastBoard),
		),
	)

	return &types.MsgPlayMoveResponse{
		CapturedX: int32(captured.X),
		CapturedY: int32(captured.Y),
		Winner:    rules.PieceStrings[game.Winner()],
	}, nil
}

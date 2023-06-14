package keeper

import (
	"context"
	"strconv"

	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the new index of the game
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("system info not found")
	}
	newIdx := strconv.FormatUint(systemInfo.NextId, 10)

	// create the new game object to store
	newGame := rules.New()
	storedGame := types.StoredGame{
		Index:       newIdx,
		Board:       newGame.String(),
		Turn:        rules.PieceStrings[newGame.Turn],
		Black:       msg.Black,
		Red:         msg.Red,
		Winner:      rules.PieceStrings[rules.NO_PLAYER],
		Deadline:    types.FormatDeadline(types.GetNextDeadline(ctx)),
		MoveCount:   0,
		BeforeIndex: types.NoFifoIndex,
		AfterIndex:  types.NoFifoIndex,
		Wager:       msg.Wager,
	}
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}

	// send the new game to the tail because it is freshly created
	k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)

	// store the new game
	k.Keeper.SetStoredGame(ctx, storedGame)

	// update checkers state
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, newIdx),
			sdk.NewAttribute(types.GameCreatedEventBlack, msg.Black),
			sdk.NewAttribute(types.GameCreatedEventRed, msg.Red),
			sdk.NewAttribute(types.GameCreatedEventWager, strconv.FormatUint(msg.Wager, 10)),
		),
	)

	return &types.MsgCreateGameResponse{
		GameIndex: newIdx,
	}, nil
}

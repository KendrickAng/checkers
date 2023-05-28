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
	storedGame := types.StoredGame {
		Index: newIdx,
		Board: newGame.String(),
		Turn: rules.PieceStrings[newGame.Turn],
		Black: msg.Black,
		Red: msg.Red,
	}
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	
	// store the new game
	k.Keeper.SetStoredGame(ctx, storedGame)

	// update checkers state
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	return &types.MsgCreateGameResponse{
		GameIndex: newIdx,
	}, nil
}

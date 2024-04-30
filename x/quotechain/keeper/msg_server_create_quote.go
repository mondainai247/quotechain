package keeper

import (
	"context"

	"quotechain/x/quotechain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateQuote(goCtx context.Context, msg *types.MsgCreateQuote) (*types.MsgCreateQuoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var quote = types.Quote{
            Creator: msg.Creator,
            Phrase: msg.Phrase,
            Author: msg.Author,
        }
        id := k.AppendQuote(
           ctx,
           quote,
          )

        return &types.MsgCreateQuoteResponse{
        Id: id,
        }, nil
}

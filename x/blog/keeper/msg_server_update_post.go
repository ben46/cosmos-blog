package keeper

import (
	"blog/x/blog/types"
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var post = types.Post{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}
	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d not found", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	k.SetPost(ctx, post)
	return &types.MsgUpdatePostResponse{}, nil
}

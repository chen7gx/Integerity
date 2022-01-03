package keeper

import (
	"context"

	"github.com/chen7gx/integrity/x/integrity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateHash(goCtx context.Context, msg *types.MsgCreateHash) (*types.MsgCreateHashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var datahash = types.Datahash{
		Creator: msg.Creator,
		Details: msg.Details,
		Hash:    msg.Hash,
	}
	// Add a post to the store and get back the ID
	id := k.AppendDatahash(ctx, datahash)

	// Return the ID of the post

	txgas := sdk.NewInt(1)
<<<<<<< HEAD
	//生成代币的数量为1
	coin := sdk.NewCoin("aphoton", txgas)
	//定义代币：1apoton
=======
	coin := sdk.NewCoin("stake", txgas)
>>>>>>> 69f7f424a2aadc589c02fadadc67130bc372c696
	coins := sdk.NewCoins(coin)

	err := k.MintCoinsForHash(ctx, coins)
	if err != nil {
		panic(err)
	}
	//使用mint模块实际生成代币

	creatorAddr, err := sdk.AccAddressFromBech32(datahash.Creator)
	if err != nil {
		panic(err)
	}
	//获取发起交易的节点的地址

	err = k.SendCoinsFromMintModuleToAccount(ctx, coins, creatorAddr)
	if err != nil {
		panic(err)
	}
	//将代币转移给发起交易的节点

	return &types.MsgCreateHashResponse{Id: id}, nil
}

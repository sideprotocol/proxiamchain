package keeper_test

import (
	"context"
	"testing"

	keepertest "bobchain/testutil/keeper"
	"bobchain/x/bobchain/keeper"
	"bobchain/x/bobchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BobchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

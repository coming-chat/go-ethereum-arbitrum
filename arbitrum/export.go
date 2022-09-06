package arbitrum

import (
	"context"

	"github.com/coming-chat/go-ethereum-arbitrum/common/hexutil"
	"github.com/coming-chat/go-ethereum-arbitrum/core"
	"github.com/coming-chat/go-ethereum-arbitrum/internal/ethapi"
	"github.com/coming-chat/go-ethereum-arbitrum/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}

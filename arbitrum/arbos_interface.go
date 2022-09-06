package arbitrum

import (
	"context"

	"github.com/coming-chat/go-ethereum-arbitrum/core"
	"github.com/coming-chat/go-ethereum-arbitrum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}

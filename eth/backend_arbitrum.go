package eth

import (
	"github.com/coming-chat/go-ethereum-arbitrum/core"
	"github.com/coming-chat/go-ethereum-arbitrum/core/state"
	"github.com/coming-chat/go-ethereum-arbitrum/core/types"
	"github.com/coming-chat/go-ethereum-arbitrum/core/vm"
	"github.com/coming-chat/go-ethereum-arbitrum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, error) {
	return eth.stateAtTransaction(block, txIndex, reexec)
}

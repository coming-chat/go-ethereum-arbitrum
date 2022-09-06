package vm

import "github.com/coming-chat/go-ethereum-arbitrum/common"

var (
	PrecompiledContractsArbitrum = make(map[common.Address]PrecompiledContract)
	PrecompiledAddressesArbitrum []common.Address
)

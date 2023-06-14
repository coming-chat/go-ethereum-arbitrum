package types

import (
	"bytes"

	"github.com/coming-chat/go-ethereum-arbitrum/common"
	"github.com/coming-chat/go-ethereum-arbitrum/rlp"
)

type ZksyncUnsignTxData struct {
	LegacyTx
	HashOverride common.Hash // Hash cannot be locally computed from other fields
	TxType       byte
}

func (tx *ZksyncUnsignTxData) copy() TxData {
	legacyCopy := tx.LegacyTx.copy().(*LegacyTx)
	return &OptimismLegacyTxData{
		LegacyTx:     *legacyCopy,
		HashOverride: tx.HashOverride,
	}
}

func (tx *ZksyncUnsignTxData) txType() byte { return tx.TxType }

func (tx *ZksyncUnsignTxData) EncodeOnlyLegacyInto(w *bytes.Buffer) {
	rlp.Encode(w, tx.LegacyTx)
}

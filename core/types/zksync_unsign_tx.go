package types

import (
	"bytes"
	"errors"

	"github.com/coming-chat/go-ethereum-arbitrum/common"
	"github.com/coming-chat/go-ethereum-arbitrum/rlp"
)

type ZksyncUnsignTxData struct {
	LegacyTx
	HashOverride common.Hash // Hash cannot be locally computed from other fields
}

func NewZksyncUnsignTx(origTx *Transaction, hashOverride common.Hash, effectiveGas uint64, l1Block uint64) (*Transaction, error) {
	if origTx.Type() != LegacyTxType {
		return nil, errors.New("attempt to op-wrap non-legacy transaction")
	}
	legacyPtr := origTx.GetInner().(*LegacyTx)
	inner := OptimismLegacyTxData{
		LegacyTx:     *legacyPtr,
		HashOverride: hashOverride,
	}
	return NewTx(&inner), nil
}

func (tx *ZksyncUnsignTxData) copy() TxData {
	legacyCopy := tx.LegacyTx.copy().(*LegacyTx)
	return &OptimismLegacyTxData{
		LegacyTx:     *legacyCopy,
		HashOverride: tx.HashOverride,
	}
}

func (tx *ZksyncUnsignTxData) txType() byte { return OptimismLegacyTxType }

func (tx *ZksyncUnsignTxData) EncodeOnlyLegacyInto(w *bytes.Buffer) {
	rlp.Encode(w, tx.LegacyTx)
}

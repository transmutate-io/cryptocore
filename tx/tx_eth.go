package tx

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/transmutate-io/cryptocore/types"
)

var (
	_ Tx           = (*TxETH)(nil)
	_ TxStateBased = (*TxETH)(nil)
)

type TxETH struct{ Tx *ethtypes.Transaction }

func (tx *TxETH) Hash() types.Bytes { return types.Bytes(tx.Tx.Hash().Bytes()) }

func (tx *TxETH) ID() types.Bytes { return tx.Hash() }

func (tx *TxETH) To() string { return tx.Tx.To().Hex() }

func (tx *TxETH) Value() types.Amount { return types.NewAmountBig(tx.Tx.Value(), 18) }

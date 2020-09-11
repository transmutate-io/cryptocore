package block

import (
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

type (
	ForwardBlockNavigator interface {
		NextBlockHash() types.Bytes
	}

	BackwardBlockNavigator interface {
		PreviousBlockHash() types.Bytes
	}

	ConfirmationCounter interface {
		Confirmations() int
	}

	TransactionsLister interface {
		TransactionsHashes() []types.Bytes
	}

	TransactionsFetcher interface {
		Transactions() []tx.Tx
	}

	Block interface {
		Hash() types.Bytes
		Height() int
		Time() types.UnixTime
	}
)

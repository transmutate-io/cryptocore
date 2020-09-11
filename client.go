package cryptocore

import (
	"github.com/transmutate-io/cryptocore/block"
	"github.com/transmutate-io/cryptocore/tx"
	"github.com/transmutate-io/cryptocore/types"
)

type (
	BlockFunc       = func() (block.Block, error)
	TransactionFunc = func() (tx.Tx, error)
	CloseFunc       = func()

	BlockGenerator interface {
		Generate(nBlocks int) ([]types.Bytes, error)
	}

	TargetedBlockGenerator interface {
		GenerateToAddress(nBlocks int, addr string) ([]types.Bytes, error)
	}

	AddressGenerator interface {
		NewAddress() (string, error)
	}

	Sender interface {
		SendToAddress(addr string, value types.Amount) (types.Bytes, error)
	}

	TransactionSender interface {
		SendTransaction(t tx.Tx) (types.Bytes, error)
	}

	RawTransactionSender interface {
		SendRawTransaction(tx types.Bytes) (types.Bytes, error)
	}

	Balancer interface {
		Balance(minConf int64) (types.Amount, error)
	}

	TargetedBalancer interface {
		BalanceOf(addr string, minConf int64) (types.Amount, error)
	}

	AddressLister interface {
		AvailableAddresses() ([]string, error)
	}

	Client interface {
		BlockCount() (uint64, error)
		BlockHash(height uint64) (types.Bytes, error)
		RawBlock(hash types.Bytes) (types.Bytes, error)
		Block(hash types.Bytes) (block.Block, error)
		RawTransaction(hash types.Bytes) (types.Bytes, error)
		Transaction(hash types.Bytes) (tx.Tx, error)
	}
	TLSConfig struct {
		ClientCertificate string
		ClientKey         string
		CA                string
		SkipVerify        bool
	}
)

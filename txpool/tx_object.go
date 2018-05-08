package txpool

import (
	"math/big"

	"github.com/vechain/thor/tx"
)

type objectStatus int

const (
	Pending objectStatus = iota
	Queued
)

//txObject wrap transaction
type txObject struct {
	tx           *tx.Transaction
	status       objectStatus
	overallGP    *big.Int
	creationTime int64
	deleted      bool
}

type txObjects []*txObject

func (txObjs txObjects) parseTxs() []*tx.Transaction {
	txs := make(tx.Transactions, 0, len(txObjs))
	for _, obj := range txObjs {
		if !obj.deleted {
			txs = append(txs, obj.tx)
		}
	}
	return txs
}

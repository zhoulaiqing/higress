package core

import (
	"strings"
	"sync"
)

type WAF struct {
	txPool sync.Pool

	RuleEngine RuleEngine
}

func (w *WAF) NewTransaction() *Transaction {
	return w.NewTransactionWithId(RandomString(16))
}

func (w *WAF) NewTransactionWithId(id string) *Transaction {
	if len(strings.TrimSpace(id)) == 0 {
		id = RandomString(16)
	}
	return w.newTransactionWithId(id)
}

func (w *WAF) newTransactionWithId(id string) *Transaction {
	tx := w.txPool.Get().(*Transaction)
	tx.id = id

	return tx
}

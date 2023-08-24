package core

type Rule interface {
	Id() string
	Evaluate(tx Transaction) bool
}

const (
	PASS  = 1
	BLOCK = 0
	DENY  = -1
)

package core

type Rule interface {
	Id() string
	Phase() int
	Evaluate(tx Transaction) bool
}

const (
	PASS  = 1
	BLOCK = 0
	DENY  = -1
)

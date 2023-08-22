package core

type Rule interface {
	Id() string
	Evaluate(tx Transaction) bool
}

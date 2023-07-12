package main

import "fmt"

type IPInt struct {
	integers []uint64
}

const (
	LessThan    int = -1
	GreaterThan int = 1
	Equals      int = 0
)

func (u *IPInt) Cmp(other *IPInt) int {
	if len(u.integers) < len(other.integers) {
		return LessThan
	}
	if len(u.integers) > len(other.integers) {
		return GreaterThan
	}

	for i := 0; i < len(u.integers); i++ {
		if u.integers[i] < other.integers[i] {
			return LessThan
		}
		if u.integers[i] > other.integers[i] {
			return GreaterThan
		}
	}

	return Equals
}

func (u *IPInt) LessThan(other *IPInt) bool {
	return u.Cmp(other) == LessThan
}

func (u *IPInt) LessThanOrEqual(other *IPInt) bool {
	return u.Cmp(other) == LessThan || u.Cmp(other) == Equals
}

func (u *IPInt) GreaterThan(other *IPInt) bool {
	return u.Cmp(other) == GreaterThan
}

func (u *IPInt) GreaterThanOrEqual(other *IPInt) bool {
	return u.Cmp(other) == GreaterThan || u.Cmp(other) == Equals
}

func (u *IPInt) EqualTo(other *IPInt) bool {
	return u.Cmp(other) == Equals
}

func (u *IPInt) BitAnd(other *IPInt) (*IPInt, error) {
	if len(u.integers) != len(other.integers) {
		return nil, fmt.Errorf("BitAnd Error: Not the same length.")
	}

	for i := 0; i < len(u.integers); i++ {
		u.integers[i] &= other.integers[i]
	}

	return u, nil
}

func (u *IPInt) BitOr(other *IPInt) (*IPInt, error) {
	if len(u.integers) != len(other.integers) {
		return nil, fmt.Errorf("BitOr Error: Not the same length.")
	}

	for i := 0; i < len(u.integers); i++ {
		u.integers[i] |= other.integers[i]
	}

	return u, nil
}

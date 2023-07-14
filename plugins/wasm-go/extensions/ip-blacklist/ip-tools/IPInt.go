package ip_tools

import (
	"fmt"
	"strconv"
	"strings"
)

type IPInt struct {
	integers []uint32
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

	var integers []uint32

	for i := 0; i < len(u.integers); i++ {
		integers = append(integers, u.integers[i]&other.integers[i])
	}

	return &IPInt{integers: integers}, nil
}

func (u *IPInt) BitOr(other *IPInt) (*IPInt, error) {
	if len(u.integers) != len(other.integers) {
		return nil, fmt.Errorf("BitOr Error: Not the same length.")
	}

	var integers []uint32

	for i := 0; i < len(u.integers); i++ {
		integers = append(integers, u.integers[i]|other.integers[i])
	}

	return &IPInt{integers}, nil
}

func (u *IPInt) BitInverse() (*IPInt, error) {
	var integers []uint32

	for i := 0; i < len(u.integers); i++ {
		integers = append(integers, ^u.integers[i])
	}

	return &IPInt{integers}, nil
}

func (u *IPInt) String() string {
	strArr := make([]string, len(u.integers))

	for i, val := range u.integers {
		strArr[i] = strconv.FormatUint(uint64(val), 10)
	}

	return strings.Join(strArr, ",")
}

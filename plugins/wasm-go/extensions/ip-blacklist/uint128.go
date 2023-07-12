package main

import (
	"fmt"
	"math/big"
)

type uint128 struct {
	lo uint64
	hi uint64
}

func newUint128(lo, hi uint64) uint128 {
	return uint128{lo, hi}
}

func (u uint128) String() string {
	return fmt.Sprintf("%x%016x", u.hi, u.lo)
}

func addUint128(a, b uint128) uint128 {
	lo := a.lo + b.lo
	hi := a.hi + b.hi
	if lo < a.lo {
		hi++
	}
	return uint128{lo, hi}
}

func subUint128(a, b uint128) uint128 {
	lo := a.lo - b.lo
	hi := a.hi - b.hi
	if lo > a.lo {
		hi--
	}
	return uint128{lo, hi}
}

func mulUint128(a, b uint128) uint128 {
	// Convert uint128 to big.Int
	aBigInt := big.NewInt(0).SetUint64(a.lo)
	aBigInt.Lsh(aBigInt, 64)
	aBigInt.Add(aBigInt, big.NewInt(0).SetUint64(a.hi))

	bBigInt := big.NewInt(0).SetUint64(b.lo)
	bBigInt.Lsh(bBigInt, 64)
	bBigInt.Add(bBigInt, big.NewInt(0).SetUint64(b.hi))

	// Perform multiplication using big.Int
	cBigInt := big.NewInt(0).Mul(aBigInt, bBigInt)

	// Convert big.Int back to uint128
	c := uint128{}
	c.lo = cBigInt.Uint64()
	cBigInt.Rsh(cBigInt, 64)
	c.hi = cBigInt.Uint64()

	return c
}

func (u uint128) LessThan(other uint128) bool {
	if u.hi < other.hi {
		return true
	} else if u.hi == other.hi {
		return u.lo < other.lo
	}
	return false
}

func (u uint128) GreaterThan(other uint128) bool {
	if u.hi > other.hi {
		return true
	} else if u.hi == other.hi {
		return u.lo > other.lo
	}
	return false
}

func (u uint128) EqualTo(other uint128) bool {
	return u.hi == other.hi && u.lo == other.lo
}

//func main() {
//	a := newUint128(0xffffffffffffffff, 0xffffffffffffffff)
//	b := newUint128(0x0000000000000001, 0x0000000000000001)
//
//	// Addition
//	c := addUint128(a, b)
//	fmt.Println("Addition:", c)
//
//	// Subtraction
//	d := subUint128(a, b)
//	fmt.Println("Subtraction:", d)
//
//	// Multiplication
//	e := mulUint128(a, b)
//	fmt.Println("Multiplication:", e)
//}

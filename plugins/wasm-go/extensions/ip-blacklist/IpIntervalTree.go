package main

import "math/big"

type IPInterval struct {
	Start *big.Int
	End   *big.Int
}

type IPIntervalNode struct {
	Interval    IPInterval
	MaxEnd      *big.Int
	Left, Right *IPIntervalNode
}

func NewIPIntervalNode(interval IPInterval) *IPIntervalNode {
	return &IPIntervalNode{
		Interval: interval,
		MaxEnd:   interval.End,
		Left:     nil,
		Right:    nil,
	}
}

func IPIntervalTreeInsert(root *IPIntervalNode, interval IPInterval) *IPIntervalNode {
	if root == nil {
		return NewIPIntervalNode(interval)
	}

	if interval.Start.Cmp(root.Interval.Start) < 0 {
		root.Left = IPIntervalTreeInsert(root.Left, interval)
	} else {
		root.Right = IPIntervalTreeInsert(root.Right, interval)
	}

	if root.MaxEnd.Cmp(interval.End) < 0 {
		root.MaxEnd = interval.End
	}

	return root
}

func IPSearch(root *IPIntervalNode, ipInt *big.Int) *IPInterval {
	if root == nil {
		return nil
	}

	if ipInt.Cmp(root.Interval.Start) >= 0 && ipInt.Cmp(root.Interval.End) <= 0 {
		return &root.Interval
	}

	if root.Left != nil && root.Left.MaxEnd.Cmp(ipInt) >= 0 {
		return IPSearch(root.Left, ipInt)
	}

	return IPSearch(root.Right, ipInt)
}

package main

type Interval struct {
	Start uint64
	End   uint64
}

type Node struct {
	Interval    Interval
	MaxEnd      uint64
	Left, Right *Node
}

func NewNode(interval Interval) *Node {
	return &Node{
		Interval: interval,
		MaxEnd:   interval.End,
		Left:     nil,
		Right:    nil,
	}
}

func Insert(root *Node, interval Interval) *Node {
	if root == nil {
		return NewNode(interval)
	}

	if interval.Start < root.Interval.Start {
		root.Left = Insert(root.Left, interval)
	} else {
		root.Right = Insert(root.Right, interval)
	}

	if root.MaxEnd < interval.End {
		root.MaxEnd = interval.End
	}

	return root
}

func Search(root *Node, point uint64) *Interval {
	if root == nil {
		return nil
	}

	if point >= root.Interval.Start && point <= root.Interval.End {
		return &root.Interval
	}

	if root.Left != nil && root.Left.MaxEnd >= point {
		return Search(root.Left, point)
	}

	return Search(root.Right, point)
}

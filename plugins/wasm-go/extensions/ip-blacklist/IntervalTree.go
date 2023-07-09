package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

type Node struct {
	Interval    Interval
	MaxEnd      int
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

func Search(root *Node, point int) *Interval {
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

func main() {
	root := Insert(nil, Interval{Start: 15, End: 20})
	root = Insert(root, Interval{Start: 10, End: 30})
	root = Insert(root, Interval{Start: 17, End: 19})
	root = Insert(root, Interval{Start: 5, End: 20})
	root = Insert(root, Interval{Start: 12, End: 15})
	root = Insert(root, Interval{Start: 30, End: 40})

	fmt.Println("Search results:")
	fmt.Println("Search(14):", Search(root, 14))
	fmt.Println("Search(18):", Search(root, 18))
	fmt.Println("Search(25):", Search(root, 25))
}

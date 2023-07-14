package ip_tools

type IPInterval struct {
	Start *IPInt
	End   *IPInt
}

func (i *IPInterval) String() string {
	return "(Start: " + i.Start.String() + ", End: " + i.End.String() + ")"
}

type IPIntervalNode struct {
	Interval    IPInterval
	MaxEnd      *IPInt
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
	//fmt.Println("Interval: ", interval.String())

	if root == nil {
		return NewIPIntervalNode(interval)
	}

	// 优化：如果待插入区间 interval 完全被 root 对应的区间包含，则不做任何操作
	if interval.Start.GreaterThanOrEqual(root.Interval.Start) && interval.End.LessThanOrEqual(root.Interval.End) {
		return root
	}

	if interval.Start.Cmp(root.Interval.Start) < 0 {
		root.Left = IPIntervalTreeInsert(root.Left, interval)
	} else {
		if interval.Start.LessThanOrEqual(root.Interval.End) {
			// 优化，如果待插入区间在 root 区间右侧且与 root 区间有重叠，则直接合并
			root.Interval.End = interval.End
		} else {
			root.Right = IPIntervalTreeInsert(root.Right, interval)
		}
	}

	if root.MaxEnd.Cmp(interval.End) < 0 {
		root.MaxEnd = interval.End
	}

	return root
}

func IPSearch(root *IPIntervalNode, ipInt *IPInt) *IPInterval {
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

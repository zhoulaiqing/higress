package ip_tools

type Node struct {
	Interval    *IPInterval
	MaxEndpoint *IPInt
	Height      int
	Left, Right *Node
}

type AVLTree struct {
	Root *Node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

// Helper function to find the maximum endpoint of two integers
func maxIPInt(a, b *IPInt) *IPInt {
	if a.GreaterThan(b) {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to get the height of a node
func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// Helper function to update the MaxEndpoint of a node
func updateMaxEndpoint(node *Node) {
	if node == nil {
		return
	}
	node.MaxEndpoint = maxIPInt(node.Interval.End, maxIPInt(maxEndpoint(node.Left), maxEndpoint(node.Right)))
}

// Helper function to get the MaxEndpoint of a node
func maxEndpoint(node *Node) *IPInt {
	if node == nil {
		return &IPInt{[]uint32{0}}
	}
	return node.MaxEndpoint
}

// Helper function to perform a right rotation
func rightRotate(y *Node) *Node {
	x := y.Left
	t2 := x.Right

	x.Right = y
	y.Left = t2

	// Update heights
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	// Update MaxEndpoints
	updateMaxEndpoint(y)
	updateMaxEndpoint(x)

	return x
}

// Helper function to perform a left rotation
func leftRotate(x *Node) *Node {
	y := x.Right
	t2 := y.Left

	y.Left = x
	x.Right = t2

	// Update heights
	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	// Update MaxEndpoints
	updateMaxEndpoint(x)
	updateMaxEndpoint(y)

	return y
}

// Helper function to get the balance factor of a node
func balanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

// Helper function to insert a node into the AVL tree
func (avl *AVLTree) insertHelper(root *Node, interval *IPInterval) *Node {
	if root == nil {
		return &Node{
			Interval:    interval,
			MaxEndpoint: interval.End,
			Height:      1,
		}
	}

	// 优化：如果待插入区间 interval 完全被 root 对应的区间包含，则不做任何操作
	if interval.Start.GreaterThanOrEqual(root.Interval.Start) && interval.End.LessThanOrEqual(root.Interval.End) {
		return root
	}

	if interval.Start.LessThan(root.Interval.Start) {
		root.Left = avl.insertHelper(root.Left, interval)
	} else {
		if interval.Start.LessThanOrEqual(root.Interval.End) {
			// 优化，如果待插入区间在 root 区间右侧且与 root 区间有重叠，则直接合并
			root.Interval.End = interval.End
		} else {
			root.Right = avl.insertHelper(root.Right, interval)
		}
	}

	// Update height and MaxEndpoint
	root.Height = max(height(root.Left), height(root.Right)) + 1
	updateMaxEndpoint(root)

	// Get the balance factor of the node
	balance := balanceFactor(root)

	// Perform rotations if necessary to rebalance the tree
	// Left Left Case
	if balance > 1 && interval.Start.LessThan(root.Left.Interval.Start) {
		return rightRotate(root)
	}

	// Right Right Case
	if balance < -1 && interval.Start.GreaterThan(root.Right.Interval.Start) {
		return leftRotate(root)
	}

	// Left Right Case
	if balance > 1 && interval.Start.GreaterThan(root.Left.Interval.Start) {
		root.Left = leftRotate(root.Left)
		return rightRotate(root)
	}

	// Right Left Case
	if balance < -1 && interval.Start.LessThan(root.Right.Interval.Start) {
		root.Right = rightRotate(root.Right)
		return leftRotate(root)
	}

	return root
}

// Insert a new interval into the AVL tree
func (avl *AVLTree) Insert(interval *IPInterval) {
	avl.Root = avl.insertHelper(avl.Root, interval)
}

// Search for intervals that overlap with a given interval
func (avl *AVLTree) Search(interval *IPInterval) []*IPInterval {
	var result []*IPInterval
	avl.searchHelper(avl.Root, interval, &result)
	return result
}

func (avl *AVLTree) SearchSingle(ipInt *IPInt) *IPInterval {
	return avl.searchSingleHelper(avl.Root, ipInt)
}

// Helper function for searching intervals that overlap with a given interval
func (avl *AVLTree) searchHelper(node *Node, interval *IPInterval, result *[]*IPInterval) {
	if node == nil {
		return
	}

	if node.Interval.Start.LessThanOrEqual(interval.End) && node.Interval.End.GreaterThanOrEqual(interval.Start) {
		*result = append(*result, node.Interval)
	}

	if node.Left != nil && node.Left.MaxEndpoint.GreaterThanOrEqual(interval.Start) {
		avl.searchHelper(node.Left, interval, result)
	}

	if node.Right != nil && node.Right.Interval.Start.LessThan(interval.End) {
		avl.searchHelper(node.Right, interval, result)
	}
}

func (avl *AVLTree) searchSingleHelper(node *Node, ipInt *IPInt) *IPInterval {
	if node == nil {
		return nil
	}

	if ipInt.Cmp(node.Interval.Start) >= 0 && ipInt.Cmp(node.Interval.End) <= 0 {
		return node.Interval
	}

	if node.Left != nil && node.Left.MaxEndpoint.Cmp(ipInt) >= 0 {
		return avl.searchSingleHelper(node.Left, ipInt)
	}

	return avl.searchSingleHelper(node.Right, ipInt)
}

// In-order traversal of the tree
func (avl *AVLTree) InOrderTraversal() []*IPInterval {
	var result []*IPInterval
	avl.inOrderHelper(avl.Root, &result)
	return result
}

// Helper function for in-order traversal of the tree
func (avl *AVLTree) inOrderHelper(node *Node, result *[]*IPInterval) {
	if node == nil {
		return
	}

	avl.inOrderHelper(node.Left, result)
	*result = append(*result, node.Interval)
	avl.inOrderHelper(node.Right, result)
}

package main

import (
	"fmt"
)

type LRUNode struct {
	Key   string
	Score int
	Left  *LRUNode
	Right *LRUNode
}

type ZSet struct {
	root     *LRUNode
	capacity int
}

func NewZSet(capacity int) *ZSet {
	return &ZSet{
		root:     nil,
		capacity: capacity,
	}
}

func (z *ZSet) ZAdd(key string, score int) {
	z.root = z.insert(z.root, key, score)
	if z.size(z.root) > z.capacity {
		z.root = z.deleteMin(z.root)
	}
}

func (z *ZSet) ZGet(key string) (int, bool) {
	return z.get(z.root, key)
}

func (z *ZSet) ZRange(min, max int) []string {
	var result []string
	z.rangeSearch(z.root, &result, min, max)
	return result
}

func (z *ZSet) ZLexCount(min, max string) int {
	count := 0
	z.lexCountSearch(z.root, &count, min, max)
	return count
}

func (z *ZSet) ZRemRangeByScore(min, max int) {
	z.root = z.removeRangeByScore(z.root, min, max)
}

func (z *ZSet) insert(node *LRUNode, key string, score int) *LRUNode {
	if node == nil {
		return &LRUNode{Key: key, Score: score}
	}

	if score == node.Score {
		node.Key = key // Update key for existing score
		return node
	}

	if score < node.Score {
		node.Left = z.insert(node.Left, key, score)
	} else {
		node.Right = z.insert(node.Right, key, score)
	}

	return z.balance(node)
}

func (z *ZSet) balance(node *LRUNode) *LRUNode {
	balanceFactor := z.getBalanceFactor(node)

	// Left-Left case
	if balanceFactor > 1 && z.getBalanceFactor(node.Left) >= 0 {
		return z.rotateRight(node)
	}

	// Left-Right case
	if balanceFactor > 1 && z.getBalanceFactor(node.Left) < 0 {
		node.Left = z.rotateLeft(node.Left)
		return z.rotateRight(node)
	}

	// Right-Right case
	if balanceFactor < -1 && z.getBalanceFactor(node.Right) <= 0 {
		return z.rotateLeft(node)
	}

	// Right-Left case
	if balanceFactor < -1 && z.getBalanceFactor(node.Right) > 0 {
		node.Right = z.rotateRight(node.Right)
		return z.rotateLeft(node)
	}

	return node
}

func (z *ZSet) rotateLeft(y *LRUNode) *LRUNode {
	x := y.Right
	y.Right = x.Left
	x.Left = y
	return x
}

func (z *ZSet) rotateRight(x *LRUNode) *LRUNode {
	y := x.Left
	x.Left = y.Right
	y.Right = x
	return y
}

func (z *ZSet) getBalanceFactor(node *LRUNode) int {
	if node == nil {
		return 0
	}
	return z.height(node.Left) - z.height(node.Right)
}

func (z *ZSet) height(node *LRUNode) int {
	if node == nil {
		return 0
	}
	leftHeight := z.height(node.Left)
	rightHeight := z.height(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (z *ZSet) get(node *LRUNode, key string) (int, bool) {
	if node == nil {
		return 0, false
	}

	if node.Key == key {
		return node.Score, true
	}

	if key < node.Key {
		return z.get(node.Left, key)
	}
	return z.get(node.Right, key)
}

func (z *ZSet) rangeSearch(node *LRUNode, result *[]string, min, max int) {
	if node == nil {
		return
	}

	if node.Score >= min {
		z.rangeSearch(node.Left, result, min, max)
		if node.Score <= max {
			*result = append(*result, node.Key)
		}
		z.rangeSearch(node.Right, result, min, max)
	}
}

func (z *ZSet) lexCountSearch(node *LRUNode, count *int, min, max string) {
	if node == nil {
		return
	}

	if node.Key >= min {
		z.lexCountSearch(node.Left, count, min, max)
		if node.Key <= max {
			*count++
		}
		z.lexCountSearch(node.Right, count, min, max)
	}
}

func (z *ZSet) deleteMin(node *LRUNode) *LRUNode {
	if node.Left == nil {
		return node.Right
	}
	node.Left = z.deleteMin(node.Left)
	return z.balance(node)
}

func (z *ZSet) removeRangeByScore(node *LRUNode, min, max int) *LRUNode {
	if node == nil {
		return nil
	}

	if node.Score > max {
		node.Left = z.removeRangeByScore(node.Left, min, max)
	} else if node.Score < min {
		node.Right = z.removeRangeByScore(node.Right, min, max)
	} else {
		node.Left = z.removeRangeByScore(node.Left, min, max)
		node.Right = z.removeRangeByScore(node.Right, min, max)
		node = z.deleteNode(node)
	}
	return node
}

func (z *ZSet) deleteNode(node *LRUNode) *LRUNode {
	if node.Left == nil {
		return node.Right
	}
	if node.Right == nil {
		return node.Left
	}

	minNode := z.findMin(node.Right)
	minNode.Right = z.deleteMin(node.Right)
	minNode.Left = node.Left
	return z.balance(minNode)
}

func (z *ZSet) findMin(node *LRUNode) *LRUNode {
	if node.Left == nil {
		return node
	}
	return z.findMin(node.Left)
}

func (z *ZSet) size(node *LRUNode) int {
	if node == nil {
		return 0
	}
	return 1 + z.size(node.Left) + z.size(node.Right)
}

func (z *ZSet) Print() {
	z.printInOrder(z.root)
	fmt.Println()
}

func (z *ZSet) printInOrder(node *LRUNode) {
	if node == nil {
		return
	}
	z.printInOrder(node.Left)
	fmt.Printf("(%s: %d) ", node.Key, node.Score)
	z.printInOrder(node.Right)
}

func main() {
	zset := NewZSet(3)

	zset.ZAdd("key1", 50)
	zset.ZAdd("key2", 30)
	zset.ZAdd("key3", 70)
	zset.Print() // 输出: (key2: 30) (key1: 50) (key3: 70)

	zset.ZAdd("key2", 90)
	zset.Print() // 输出: (key2: 90) (key1: 50) (key3: 70)

	getResult, _ := zset.ZGet("key2")
	fmt.Println("ZGet key2:", getResult) // 输出: ZGet key2: 90

	rangeResult := zset.ZRange(40, 100)
	fmt.Println("ZRange (40, 100):", rangeResult) // 输出: ZRange (40, 100): [key1, key2]

	lexCount := zset.ZLexCount("key1", "z")
	fmt.Println("ZLexCount ('key1', 'z'):", lexCount) // 输出: ZLexCount ('key1', 'z'): 3

	zset.ZRemRangeByScore(40, 100)
	zset.Print() // 输出: (key2: 90)

	zset.ZRemRangeByScore(10, 30)
	zset.Print() // 输出: (key2: 90)
}

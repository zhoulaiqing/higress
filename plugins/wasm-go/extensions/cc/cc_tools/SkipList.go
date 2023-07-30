package cc_tools

import (
	"bytes"
	"fmt"
	"math/rand"
)

const (
	maxLevel    = 32   // SkipList最大层数
	probability = 0.25 // 插入时节点层数增加的概率
)

// Node 表示 SkipList 中的一个节点
type Node struct {
	key     int64
	forward []*Node // 每层的下一个节点指针
}

// SkipList 表示 SkipList 的结构
type SkipList struct {
	head   *Node // 头节点
	level  int   // 当前 SkipList 层数
	length int   // SkipList 长度
}

// NewNode 创建一个新的节点
func NewNode(key int64, level int) *Node {
	return &Node{
		key:     key,
		forward: make([]*Node, level),
	}
}

// NewSkipList 创建一个新的 SkipList
func NewSkipList() *SkipList {
	head := NewNode(0, maxLevel)
	return &SkipList{
		head:   head,
		level:  1,
		length: 0,
	}
}

// randomLevel 根据概率随机生成节点的层数
func randomLevel() int {
	level := 1
	for rand.Float32() < probability && level < maxLevel {
		level++
	}
	return level
}

// Insert 插入一个节点到 SkipList
func (sl *SkipList) Insert(key int64) {
	update := make([]*Node, maxLevel)
	x := sl.head

	// 找到每层的插入位置
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}

	level := randomLevel()

	// 如果新节点的层数大于当前 SkipList 的层数，更新当前层数
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	// 插入新节点
	x = NewNode(key, level)
	for i := 0; i < level; i++ {
		x.forward[i] = update[i].forward[i]
		update[i].forward[i] = x
	}

	sl.length++
}

// Find 查找指定 key 是否存在于 SkipList 中
func (sl *SkipList) Find(key int64) bool {
	x := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}

	x = x.forward[0]

	return x != nil && x.key == key
}

// RangeSearch 返回指定范围内的所有 key
func (sl *SkipList) RangeSearch(start, end int64) []int64 {
	result := []int64{}
	x := sl.head

	// 定位到 start 节点
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < start {
			x = x.forward[i]
		}
	}

	x = x.forward[0]

	// 遍历节点，直到 end 节点
	for x != nil && x.key <= end {
		result = append(result, x.key)
		x = x.forward[0]
	}

	return result
}

// RangeCount 返回指定范围内的 key 的数量
func (sl *SkipList) RangeCount(start, end int64) int {
	count := 0
	x := sl.head

	// 定位到 start 节点
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < start {
			x = x.forward[i]
		}
	}

	x = x.forward[0]

	// 遍历节点，直到 end 节点
	for x != nil && x.key <= end {
		count++
		x = x.forward[0]
	}

	return count
}

// Delete 删除指定的 key 节点
func (sl *SkipList) Delete(key int64) {
	update := make([]*Node, maxLevel)
	x := sl.head

	// 找到每层的插入位置
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}

	x = x.forward[0]

	if x != nil && x.key == key {
		// 从所有层级删除节点
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] == x {
				update[i].forward[i] = x.forward[i]
			}
		}

		// 更新 SkipList 层数
		for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
			sl.level--
		}

		sl.length--
	}
}

// Print 打印整个 SkipList
func (sl *SkipList) Print() {
	for i := sl.level - 1; i >= 0; i-- {
		x := sl.head.forward[i]
		fmt.Printf("Level %d: ", i)
		for x != nil {
			fmt.Printf("%d -> ", x.key)
			x = x.forward[i]
		}
		fmt.Println("nil")
	}
}

func (sl *SkipList) ToString() string {
	var buffer bytes.Buffer

	for i := sl.level - 1; i >= 0; i-- {
		x := sl.head.forward[i]
		buffer.WriteString(fmt.Sprintf("Level %d: ", i))
		for x != nil {
			buffer.WriteString(fmt.Sprintf("%d -> ", x.key))
			x = x.forward[i]
		}
		buffer.WriteString("nil\n")
	}

	return buffer.String()
}

// RangeDelete 删除处于指定范围内的所有节点
func (sl *SkipList) RangeDelete(start, end int64) {
	update := make([]*Node, maxLevel)
	x := sl.head

	// 找到每层的插入位置
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < start {
			x = x.forward[i]
		}
		update[i] = x
	}

	x = x.forward[0]

	// 删除范围内的所有节点
	for x != nil && x.key <= end {
		// 从所有层级删除节点
		for i := 0; i < len(x.forward); i++ {
			if update[i].forward[i] == x {
				update[i].forward[i] = x.forward[i]
			}
		}

		// 更新 SkipList 层数
		for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
			sl.level--
		}

		sl.length--

		x = x.forward[0]
	}
}

func TestZSet() {
	skipList := NewSkipList()

	// 插入示例数据
	for i := 1; i <= 20; i++ {
		skipList.Insert(int64(i))
	}

	fmt.Println("SkipList:")
	skipList.Print()

	// 查找示例
	fmt.Println("Find 10:", skipList.Find(10))
	fmt.Println("Find 25:", skipList.Find(25))

	// 范围查找示例
	fmt.Println("RangeFind 5 to 15:", skipList.RangeSearch(5, 15))

	// 范围计数示例
	fmt.Println("RangeCount 5 to 15:", skipList.RangeCount(5, 15))

	// 范围删除示例
	skipList.RangeDelete(5, 6)
	fmt.Println("SkipList after range deletion:")
	skipList.Print()
}

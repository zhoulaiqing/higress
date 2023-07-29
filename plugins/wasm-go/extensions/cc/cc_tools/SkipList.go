package cc_tools

import (
	"math"
	"math/rand"
)

const maxLevel = 16 // SkipList 的最大层数

type node struct {
	content int64
	forward []*node
}

type SkipList struct {
	head   *node
	level  int
	length int
}

// 创建一个新的节点
func createNode(content int64, level int) *node {
	return &node{
		content: content,
		forward: make([]*node, level),
	}
}

// NewSkipList 创建一个新的 SkipList
func NewSkipList() *SkipList {
	head := createNode(math.MinInt64, maxLevel)
	return &SkipList{
		head:   head,
		level:  1,
		length: 0,
	}
}

// 获取一个随机的层数，用于插入节点时的决策
func randomLevel() int {
	level := 1
	for rand.Float64() < 0.5 && level < maxLevel {
		level++
	}
	return level
}

// Insert 插入节点到 SkipList 中
func (sl *SkipList) Insert(content int64) {
	update := make([]*node, maxLevel)
	current := sl.head

	// 寻找每层插入位置
	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].content < content {
			current = current.forward[i]
		}
		update[i] = current
	}

	level := randomLevel()

	// 如果插入的层数大于当前 SkipList 层数，更新头节点的层数
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	// 创建新节点，并进行连接
	newNode := createNode(content, level)
	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}

	sl.length++
}

// 查找指定 content 的节点
func (sl *SkipList) searchNode(content int64) *node {
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].content < content {
			current = current.forward[i]
		}
	}

	// 找到 content 对应的节点或者该节点不存在
	if current.forward[0] != nil && current.forward[0].content == content {
		return current.forward[0]
	}

	return nil
}

// RangeSearch 查找在范围 [start, end] 内的节点，并返回它们的值
func (sl *SkipList) RangeSearch(start, end int64) []int64 {
	var result []int64
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].content < start {
			current = current.forward[i]
		}
	}

	current = current.forward[0]

	for current != nil && current.content <= end {
		result = append(result, current.content)
		current = current.forward[0]
	}

	return result
}

// RangeCount 计算在范围 [start, end] 内的节点个数
func (sl *SkipList) RangeCount(start, end int64) int {
	count := 0
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].content < start {
			current = current.forward[i]
		}
	}

	current = current.forward[0]

	for current != nil && current.content <= end {
		count++
		current = current.forward[0]
	}

	return count
}

// Delete 删除指定 content 的节点
func (sl *SkipList) Delete(content int64) bool {
	update := make([]*node, maxLevel)
	current := sl.head

	// 寻找每层删除位置
	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].content < content {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]

	if current != nil && current.content == content {
		// 逐层删除
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != current {
				break
			}
			update[i].forward[i] = current.forward[i]
		}

		// 更新头节点的层数
		for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
			sl.level--
		}

		sl.length--
		return true
	}

	return false
}

// RangeDelete 删除范围 [start, end] 内的节点
func (sl *SkipList) RangeDelete(start, end int64) {
	update := make([]*node, maxLevel)
	current := sl.head

	// 寻找每层删除位置
	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].content < start {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]

	// 删除范围内的节点
	for current != nil && current.content <= end {
		// 逐层删除
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != current {
				break
			}
			update[i].forward[i] = current.forward[i]
		}

		current = current.forward[0]
	}

	// 更新头节点的层数
	for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
		sl.level--
	}

	// 更新长度
	sl.length -= sl.RangeCount(start, end)
}

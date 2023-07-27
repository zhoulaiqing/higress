package main

import "sort"

type Zset struct {
	capacity int
	data     map[string]*BTree
}

type BTreeNode struct {
	score int64
}

type BTree struct {
	nodes []*BTreeNode
}

// ZAdd 向 Zset 添加一个 score，按顺序加入到对应 key 的树结构中
func (z *Zset) ZAdd(key string, score int64) {
	if z.data == nil {
		z.data = make(map[string]*BTree)
	}

	tree, found := z.data[key]
	if !found {
		tree = &BTree{}
		z.data[key] = tree
	}

	// 查找插入位置
	idx := sort.Search(len(tree.nodes), func(i int) bool {
		return tree.nodes[i].score >= score
	})

	newNode := &BTreeNode{score: score}

	if idx == len(tree.nodes) {
		tree.nodes = append(tree.nodes, newNode)
	} else {
		tree.nodes = append(tree.nodes[:idx+1], tree.nodes[idx:]...)
		tree.nodes[idx] = newNode
	}
}

// ZCount 计算在 Zset 中，指定 key 在分数区间 [min, max] 的节点数量
func (z *Zset) ZCount(key string, min, max int64) int {
	tree, found := z.data[key]
	if !found {
		return 0
	}

	count := 0
	for _, node := range tree.nodes {
		if node.score >= min && node.score <= max {
			count++
		}
	}

	return count
}

// ZRemRangeByScore 移除 Zset 中给定分数区间 [min, max] 的所有节点
func (z *Zset) ZRemRangeByScore(key string, min, max int64) {
	tree, found := z.data[key]
	if !found {
		return
	}

	var newNodes []*BTreeNode
	for _, node := range tree.nodes {
		if node.score < min || node.score > max {
			newNodes = append(newNodes, node)
		}
	}

	tree.nodes = newNodes
}

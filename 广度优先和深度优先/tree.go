package tree

type Node struct {
	value    int
	children []*Node
}

type Tree struct {
	root *Node
}

// BFS breadth first search
func (t *Tree) BFS() []int {
	var ans []int

	// 待搜索队列
	var queue []*Node

	if t.root == nil {
		return nil
	}

	// 先把root加入到待处理队列
	queue = append(queue, t.root)

	for len(queue) > 0 {
		ans = append(ans, queue[0].value)

		// 把此节点的所有子节点加入到队列中
		for _, child := range queue[0].children {
			queue = append(queue, child)
		}

		queue = queue[1:]
	}

	return ans
}

// DFS Depth first search
func (t *Tree) DFS() []int {
	if t.root == nil {
		return nil
	}

	var ans []int
	return dfs(t.root, ans)
}

func dfs(root *Node, rt []int) []int {
	rt = append(rt, root.value)

	// 深度优先搜索
	for _, child := range root.children {
		rt = dfs(child, rt)
	}

	return rt
}

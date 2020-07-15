package bst

import "fmt"

// 参考 https://flaviocopes.com/golang-data-structure-binary-search-tree/

// Node 定义节点.
type Node struct {
	value int   // 因为目前Go的泛型还没有发布，所以我们这里以一个int具体类型为例
	left  *Node // 左子节点
	right *Node // 右子节点
}

// BST 是一个节点的值为int类型的二叉搜索树.
type BST struct {
	root *Node
}

// Insert 插入一个元素.
func (bst *BST) Insert(value int) {
	newNode := &Node{value, nil, nil}

	// 如果二叉树为空，那么这个节点就当作跟节点
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

// 从根节点依次比较
func insertNode(root, newNode *Node) {
	if newNode.value < root.value { // 应该放到根节点的左边
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else if newNode.value > root.value { // 应该放到根节点的右边
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}

	// 否则等于根节点
}

// Remove 删除一个元素.
func (bst *BST) Remove(value int) bool {
	_, existed := remove(bst.root, value)
	return existed
}

// 用来递归移除节点的辅助方法.
// 返回替换root的新节点，以及元素是否存在
func remove(root *Node, value int) (*Node, bool) {
	if root == nil {
		return nil, false
	}

	var existed bool
	// 从左边找
	if value < root.value {
		root.left, existed = remove(root.left, value)
		return root, existed
	}

	// 从右边找
	if value > root.value {
		root.right, existed = remove(root.right, value)
		return root, existed
	}

	// 如果此节点正是要移除的节点,那么返回此节点，同时返回之前可能需要调整.
	existed = true

	// 如果此节点没有孩子，直接返回即可
	if root.left == nil && root.right == nil {
		root = nil
		return root, existed
	}

	// 如果左子节点为空, 提升右子节点
	if root.left == nil {
		root = root.right
		return root, existed
	}
	// 如果右子节点为空, 提升左子节点
	if root.right == nil {
		root = root.left
		return root, existed
	}

	// 如果左右节点都存在,那么从左边节点找到一个最小的节点提升，这个节点肯定比左子树所有节点都大.
	// 也可以从左子树节点中找一个最大的提升，道理一样.

	smallestInRight, _ := min(root.right)
	// 提升
	root.value = smallestInRight
	// 从右边子树中移除此节点
	root.right, _ = remove(root.right, smallestInRight)

	return root, existed
}

// Search 搜索元素(检查元素是否存在)
func (bst *BST) Search(value int) bool {
	return search(bst.root, value)
}

// Min 二叉搜索树中的最小值
func (bst *BST) Min() (int, bool) {
	return min(bst.root)
}

func min(node *Node) (int, bool) {
	if node == nil {
		return 0, false
	}

	n := node
	// 从左边找
	for {
		if n.left == nil {
			return n.value, true
		}
		n = n.left
	}
}

// Max 二叉搜索树中的最大值
func (bst *BST) Max() (int, bool) {
	return max(bst.root)
}

func max(node *Node) (int, bool) {
	if node == nil {
		return 0, false
	}

	n := node
	// 从右边找
	for {
		if n.right == nil {
			return n.value, true
		}
		n = n.right
	}
}

func search(n *Node, value int) bool {
	if n == nil {
		return false
	}
	if value < n.value {
		return search(n.left, value)
	}
	if value > n.value {
		return search(n.right, value)
	}
	return true
}

// InOrderTraverse 中序遍历
func (bst *BST) InOrderTraverse(f func(int)) {
	inOrderTraverse(bst.root, f)
}
func inOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value) // 中
		inOrderTraverse(n.right, f)
	}
}

// PreOrderTraverse 前序遍历
func (bst *BST) PreOrderTraverse(f func(int)) {
	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		f(n.value) // 前
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}

// PostOrderTraverse 后序遍历
func (bst *BST) PostOrderTraverse(f func(int)) {
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value) // 后
	}
}

func (bst *BST) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.value)
		stringify(n.right, level)
	}
}

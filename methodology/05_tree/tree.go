package main

import (
	"fmt"
)

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BinarySearchTree 二叉搜索树容器
type BinarySearchTree struct {
	Root *TreeNode
}

// Insert 插入节点
// 这是一个对外暴露的包装方法
func (bst *BinarySearchTree) Insert(val int) {
	if bst.Root == nil {
		bst.Root = &TreeNode{Val: val}
	} else {
		bst.insertNode(bst.Root, val)
	}
}

// insertNode 递归插入逻辑
// 核心思想：比我小往左走，比我大往右走，走到空地就种树
func (bst *BinarySearchTree) insertNode(node *TreeNode, val int) {
	if val < node.Val {
		// 往左找
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			bst.insertNode(node.Left, val) // 递归
		}
	} else {
		// 往右找
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		} else {
			bst.insertNode(node.Right, val) // 递归
		}
	}
}

// Search 查找节点：O(log N)
func (bst *BinarySearchTree) Search(val int) bool {
	return bst.searchNode(bst.Root, val)
}

// searchNode 递归查找逻辑
func (bst *BinarySearchTree) searchNode(node *TreeNode, val int) bool {
	if node == nil {
		return false // 走到尽头也没找到
	}

	if val == node.Val {
		return true // 找到了
	} else if val < node.Val {
		return bst.searchNode(node.Left, val) // 递归左搜
	} else {
		return bst.searchNode(node.Right, val) // 递归右搜
	}
}

// InOrderTraversal 中序遍历
// 顺序：左 -> 根 -> 右
// 作用：对于BST，中序遍历的结果就是从小到大的有序序列
func (bst *BinarySearchTree) InOrderTraversal() {
	bst.inOrder(bst.Root)
	fmt.Println()
}

func (bst *BinarySearchTree) inOrder(node *TreeNode) {
	if node != nil {
		bst.inOrder(node.Left)      // 先打印左边
		fmt.Printf("%d ", node.Val) // 再打印自己
		bst.inOrder(node.Right)     // 最后打印右边
	}
}

func main() {
	bst := BinarySearchTree{}

	// 1. 构建树
	// 插入顺序：50 -> 30 -> 20 -> 40 -> 70 -> 60 -> 80
	// 这是一个典型的构建过程
	nodes := []int{50, 30, 20, 40, 70, 60, 80}
	for _, n := range nodes {
		bst.Insert(n)
	}

	fmt.Println("1. 中序遍历结果（应该是排序的）：")
	bst.InOrderTraversal() // 20 30 40 50 60 70 80

	fmt.Println("2. 查找测试：")
	fmt.Printf("Search 40: %v\n", bst.Search(40)) // true
	fmt.Printf("Search 90: %v\n", bst.Search(90)) // false
}

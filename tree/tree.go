package main

import (
	"container/list"
	"fmt"
)

var dfsnums = make([]int, 0)

/* 二叉树节点结构体 */
type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Height int // 节点高度
}

/* 构造方法 */
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Left:  nil, // 左子节点指针
		Right: nil, // 右子节点指针
		Val:   v,   // 节点值
	}
}

/* 层序遍历 */
func levelOrder(root *TreeNode) []*TreeNode {
	queue := list.New()
	queue.PushBack(root)
	// 初始化一个切片，用于保存遍历序列
	nums := make([]*TreeNode, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		nums = append(nums, node)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}

/* 前序遍历 */
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	dfsnums = append(dfsnums, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

/* 中序遍历 */
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 访问优先级：左子树 -> 根节点 -> 右子树
	inOrder(node.Left)
	dfsnums = append(dfsnums, node.Val)
	inOrder(node.Right)
}

/* 后序遍历 */
func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 访问优先级：左子树 -> 右子树 -> 根节点
	postOrder(node.Left)
	postOrder(node.Right)
	dfsnums = append(dfsnums, node.Val)
}

func main() {
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)
	// 构建引用指向（即指针）
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	//for _, v := range levelOrder(n1) {
	//	fmt.Println(v.Val)
	//}
	postOrder(n1)
	fmt.Println(dfsnums)
}

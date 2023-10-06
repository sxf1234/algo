package main

/* 二叉树节点结构体 */
type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Height int // 节点高度
}

/* 构建二叉树：分治 */
func dfsBuildTree(preorder []int, inorderMap map[int]int, preRoot, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	inRoot := inorderMap[preRoot]
	preLeftCldRoot := preRoot + 1
	preRightCldRoot := preRoot + 1 + (inRoot - left)

	leftR := dfsBuildTree(preorder, inorderMap, preLeftCldRoot, left, inRoot-1)
	rightR := dfsBuildTree(preorder, inorderMap, preRightCldRoot, inRoot+1, right)
	return &TreeNode{Val: preorder[preRoot], Left: leftR, Right: rightR}
}

/* 构建二叉树 */
func buildTree(preorder, inorder []int) *TreeNode {
	// 初始化哈希表，存储 inorder 元素到索引的映射
	inorderMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}

	root := dfsBuildTree(preorder, inorderMap, 0, 0, len(inorder)-1)
	return root
}

package main

type binarySearchTree struct {
	root *TreeNode
}

func newBinarySearchTree() *binarySearchTree {
	bst := &binarySearchTree{}
	// 初始化空树
	bst.root = nil
	return bst
}

/* 查找节点 */
func (bst *binarySearchTree) search(num int) *TreeNode {
	node := bst.root
	for node != nil {
		if node.Val == num {
			break
		} else if node.Val > num {
			node = node.Left
		} else if node.Val < num {
			node = node.Right
		}
	}
	return node
}

/* 插入节点 */
func (bst *binarySearchTree) insert(num int) {
	if bst.root == nil {
		bst.root = NewTreeNode(num)
		return
	}
	cur := bst.root
	var pre *TreeNode = nil
	for cur != nil {
		pre = cur
		if cur.Val > num {
			cur = cur.Left
		} else if cur.Val < num {
			cur = cur.Right
		} else {
			return
		}
	}
	if num > cur.Val {
		pre.Right = NewTreeNode(num)
	} else if num < cur.Val {
		pre.Left = NewTreeNode(num)
	}
}

/* 删除节点 */
func (bst *binarySearchTree) remove(num int) {
	if bst.root == nil {
		return
	}
	cur := bst.root
	var pre *TreeNode = nil
	for cur != nil {
		pre = cur
		if cur.Val > num {
			cur = cur.Left
		} else if cur.Val < num {
			cur = cur.Right
		} else {
			return
		}
	}
	// 若无待删除节点，则直接返回
	if cur == nil {
		return
	}
	if cur.Right == nil || cur.Left == nil {
		var child *TreeNode = nil
		if cur.Left == nil {
			child = cur.Right
		} else {
			child = cur.Left
		}
		//排除根节点
		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			bst.root = child
		}
		//待删除节点的子节点个数有两个
	} else {
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		bst.remove(tmp.Val)
		cur.Val = tmp.Val
	}
}

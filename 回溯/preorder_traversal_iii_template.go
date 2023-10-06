package main

//在二叉树中搜索所有值为7的节点，请返回根节点到这些节点的路径，并要求路径中不包含值为3的节点
/* 二叉树节点结构体 */
type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Height int // 节点高度
}

/* 判断当前状态是否为解 */
func isSolution(state *[]*TreeNode) bool {
	n := len(*state)
	return n != 0 && (*state)[n-1].Val == 7
}

/* 记录解 */
func recordSolution(state *[]*TreeNode, res *[][]*TreeNode) {
	*res = append(*res, *state)
}

/* 判断在当前状态下，该选择是否合法 */
func isValid(state *[]*TreeNode, choice *TreeNode) bool {
	return choice != nil && choice.Val != 3
}

/* 更新状态 */
func makeChoice(state *[]*TreeNode, choice *TreeNode) {
	*state = append(*state, choice)
}

/* 恢复状态 */
func undoChoice(state *[]*TreeNode, choice *TreeNode) {
	*state = (*state)[:len(*state)-1]
}

/* 回溯算法框架 状态 state 为节点遍历路径，选择 choices 为当前节点的左子节点和右子节点，结果 res 是路径列表。*/
func backtrack(state *[]*TreeNode, choices *[]*TreeNode, res *[][]*TreeNode) {
	// 判断是否为解
	if isSolution(state) {
		// 记录解
		recordSolution(state, res)
		// 停止继续搜索
		//return
	}
	// 遍历所有选择
	for _, choice := range *choices {
		// 剪枝：判断选择是否合法
		if isValid(state, choice) {
			// 尝试：做出选择，更新状态
			makeChoice(state, choice)
			backtrack(state, choices, res)
			// 回退：撤销选择，恢复到之前的状态
			undoChoice(state, choice)
		}
	}
}

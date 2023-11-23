package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) int // dfs返回当前节点左边路径和右边路径中的最大值
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lMax := max(0, dfs(node.Left))
		rMax := max(0, dfs(node.Right))
		ans = max(ans, lMax+rMax+node.Val) // 计算以节点为拐点的路径最大值
		return max(lMax, rMax) + node.Val
	}
	dfs(root)
	return ans
}

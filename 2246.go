package main

func longestPath(parent []int, s string) int {
	n := len(parent)
	// 建树
	fa := make([][]int, n)
	for i := 1; i < n; i++ {
		fa[parent[i]] = append(fa[parent[i]], i)
	}
	// dfs
	ans := 0
	var dfs func(cur int) int // 计算当前的节点子节点中最大路径(节点数量)
	dfs = func(cur int) int {
		// 遍历i的所有子节点，维护最大值a和次大值b
		a, b := 0, 0
		for _, i := range fa[cur] {
			t := dfs(i)
			if s[cur] == s[i] {
				t = 0
			}

			if t >= a {
				b, a = a, t
			} else if t > b {
				b = t
			}
		}
		ans = max(ans, a+b+1)
		return a + 1
	}
	dfs(0)
	return ans
}

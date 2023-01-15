package p2246

func longestPath(parent []int, s string) int {
	if len(parent) <= 1 {
		return len(parent)
	}
	var tree = make([][]int, len(parent))
	var result = 0
	for i := range parent {
		if parent[i] == -1 {
			continue
		}
		tree[parent[i]] = append(tree[parent[i]], i)
	}
	var dfs func(node int) int

	dfs = func(node int) int {
		var maxCnt1, maxCnt2 = 0, 0
		for _, child := range tree[node] {
			childCnt := dfs(child)
			result = max(result, childCnt)
			if s[node] == s[child] {
				continue
			}
			if childCnt >= maxCnt1 {
				maxCnt2 = maxCnt1
				maxCnt1 = childCnt
			} else if childCnt > maxCnt2 {
				maxCnt2 = childCnt
			}
		}
		result = max(result, maxCnt1+maxCnt2+1)
		return maxCnt1 + 1
	}
	dfs(0)
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

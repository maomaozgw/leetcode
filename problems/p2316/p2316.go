package p2316

func countPairs(n int, edges [][]int) int64 {
	var g = map[int][]int{}

	for idx := range edges {
		g[edges[idx][0]] = append(g[edges[idx][0]], edges[idx][1])
		g[edges[idx][1]] = append(g[edges[idx][1]], edges[idx][0])
	}
	var visted = make([]bool, n+1)
	var dfs func(node int)
	var vistedCnt = 0
	dfs = func(node int) {
		if visted[node] {
			return
		}
		visted[node] = true
		vistedCnt++
		for _, next := range g[node] {
			if visted[next] {
				continue
			}
			dfs(next)
		}
	}
	var cntSlice []int
	for i := 0; i < n; i++ {
		if visted[i] {
			continue
		}
		vistedCnt = 0
		dfs(i)
		cntSlice = append(cntSlice, vistedCnt)
	}
	var result = int64(0)
	var total = int64(0)
	for i := range cntSlice {
		result += int64(cntSlice[i]) * (int64(n) - total - int64(cntSlice[i]))
		total += int64(cntSlice[i])
	}
	return result
}

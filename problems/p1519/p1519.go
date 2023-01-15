package p1519

const (
	start = 'a'
)

func countSubTrees(n int, edges [][]int, labels string) []int {
	if len(edges) == 0 {
		return []int{}
	}
	var tree = map[int][]int{}
	for idx := range edges {
		tree[edges[idx][0]] = append(tree[edges[idx][0]], edges[idx][1])
		tree[edges[idx][1]] = append(tree[edges[idx][1]], edges[idx][0])
	}
	var res = make([]int, n)

	var dfs func(node, parent int) []int
	dfs = func(node, parent int) []int {
		var count = make([]int, 26)
		for _, child := range tree[node] {
			if child == parent {
				continue
			}
			childCount := dfs(child, node)
			for i := range childCount {
				count[i] += childCount[i]
			}
		}
		count[labels[node]-start] += 1
		res[node] = count[labels[node]-start]
		return count
	}
	dfs(0, -1)
	return res
}

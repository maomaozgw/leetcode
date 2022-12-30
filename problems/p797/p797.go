package p797

func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	var prefix = []int{0}
	dfs(graph, 0, len(graph)-1, prefix, &result)
	return result
}

func dfs(graph [][]int, start int, end int, prefix []int, results *[][]int) {
	for _, next := range graph[start] {
		var walked = false
		for _, val := range prefix {
			if next == val {
				walked = true
				break
			}
		}
		if walked {
			continue
		}
		var newPrefix []int
		newPrefix = append(newPrefix, prefix...)
		newPrefix = append(newPrefix, next)
		if next == end {
			*results = append(*results, newPrefix)
			continue
		}
		dfs(graph, next, end, newPrefix, results)
	}
}

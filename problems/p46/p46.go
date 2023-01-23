package p46

func permute(nums []int) [][]int {
	var result [][]int
	var dfs func(nums, path []int)
	dfs = func(remainNums, path []int) {
		if len(remainNums) == 0 {
			result = append(result, path)
			return
		}
		for i := range remainNums {
			numCopy := make([]int, len(remainNums)-1)
			pathCopy := make([]int, len(path)+1)
			copy(pathCopy, path)
			pathCopy[len(path)] = remainNums[i]
			copy(numCopy, remainNums[:i])
			copy(numCopy[i:], remainNums[i+1:])
			dfs(numCopy, pathCopy)
		}
	}
	dfs(nums, []int{})
	return result
}

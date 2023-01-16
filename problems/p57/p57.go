package p57

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	var inserted = false
	for i := range intervals {
		if inserted || intervals[i][1] < newInterval[0] {
			// 已经插入过了，或者还没开始有 overlap
			result = append(result, intervals[i])
			continue
		}
		if intervals[i][0] > newInterval[1] {
			// 当前节点的开始大于插入节点的结束(需要action)
			result = append(result, newInterval)
			inserted = true
			result = append(result, intervals[i])
			continue
		}
		// 存在 overlap
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
	}
	if !inserted {
		result = append(result, newInterval)
	}
	return result
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

package p1376

const (
	offset = 1000000
)

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var managerMap = map[int][]int{}
	for idx := range manager {
		if manager[idx] == -1 {
			continue
		}
		managerMap[manager[idx]] = append(managerMap[manager[idx]], idx)
	}
	var q []int
	q = append(q, headID)
	var result = 0
	for len(q) > 0 {
		var newQ []int
		for i := range q {
			id, cost := q[i]%offset, q[i]/offset
			result = max(result, cost)
			for _, val := range managerMap[id] {
				newQ = append(newQ, val+(cost+informTime[id])*offset)
			}
		}
		q = newQ
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

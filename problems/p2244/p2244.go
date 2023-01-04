package p2244

func minimumRounds(tasks []int) int {
	var taskCnt = map[int]int{}
	for i := range tasks {
		taskCnt[tasks[i]]++
	}
	var result = 0
	for _, val := range taskCnt {
		if val == 1 {
			return -1
		}
		result += (val + 2) / 3

	}
	return result
}

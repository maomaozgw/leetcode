package p3375

func minOperations(nums []int, k int) int {
	var min = 101
	var diffCnt = 0
	var numCnt = make([]int, min)
	for _, num := range nums {
		if numCnt[num] == 0 {
			diffCnt++
			if num < min {
				min = num
			}
		}
		numCnt[num]++
	}
	if min < k {
		return -1
	} else if min == k {
		return diffCnt - 1
	}
	return diffCnt
}

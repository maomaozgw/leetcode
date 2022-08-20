// https://leetcode.com/problems/split-array-into-consecutive-subsequences/

package p659

func isPossible(nums []int) bool {
	min := nums[0]
	max := nums[len(nums)-1]
	l := max - min + 4
	var numCount = make([]int, l)
	var needNumCount = make([]int, l)
	var lastIndex = 0

	for i := range nums {
		n := nums[i]
		lastIndex = n - min
		numCount[lastIndex]++
	}
	for i := range nums {
		n := nums[i]
		index := n - min
		if numCount[index] == 0 {
			continue
		} else if needNumCount[index] > 0 {
			needNumCount[index]--
			numCount[index]--
			needNumCount[index+1]++
		} else if numCount[index] > 0 && numCount[index+1] > 0 && numCount[index+2] > 0 {
			numCount[index]--
			numCount[index+1]--
			numCount[index+2]--
			needNumCount[index+3]++
		} else {
			return false
		}
	}
	return true
}

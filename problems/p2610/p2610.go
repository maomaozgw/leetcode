package p2610

func findMatrix(nums []int) [][]int {
	var result [][]int // 2D array
	var numIdx = make([]int, 201)
	for i := range nums {
		row := numIdx[nums[i]]
		if row >= len(result) {
			result = append(result, []int{nums[i]})
		} else {
			result[row] = append(result[row], nums[i])
		}
		numIdx[nums[i]]++
	}
	return result
}

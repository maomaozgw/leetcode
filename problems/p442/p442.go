package p442

func findDuplicates(nums []int) []int {
	var dataMap = map[int]int{}
	var result []int
	for _, num := range nums {
		if dataMap[num] == 1 {
			result = append(result, num)
			dataMap[num]++
			continue
		}
		dataMap[num] = 1
	}
	return result
}

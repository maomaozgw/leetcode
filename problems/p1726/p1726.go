package p1726

func tupleSameProduct(nums []int) int {
	var productMap = map[int]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			productMap[nums[i]*nums[j]]++
		}
	}
	var result = 0
	for _, count := range productMap {
		result += count * (count - 1) / 2
	}
	result *= 8
	return result
}

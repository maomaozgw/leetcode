package p1814

const (
	modo = 1000000007
)

func countNicePairs(nums []int) int {
	var targetMap = map[int]int{}
	var result = 0
	for i := 0; i < len(nums); i++ {
		val := nums[i] - rev(nums[i])
		result = (result + targetMap[val]) % modo
		targetMap[val] += 1
	}
	return result
}

func rev(i int) int {
	var result = 0
	for i > 0 {
		result = result*10 + i%10
		i /= 10
	}
	return result
}

package p167

func twoSum(numbers []int, target int) []int {
	var left, right = 0, len(numbers) - 1
	for left != right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

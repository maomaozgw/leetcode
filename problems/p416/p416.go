package p416

func canPartition(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if totalSum%2 != 0 {
		return false
	}

	return do(nums, totalSum/2)
}

func do(nums []int, capacity int) bool {
	dp := make([][]int, len(nums)+1)

	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, capacity+1)
	}

	for size := 1; size <= capacity; size++ {
		for idx, num := range nums {
			if num == capacity {
				return true
			}
			row := idx + 1

			if num == size || dp[row-1][size] == 1 {
				dp[row][size] = 1
			}

			if num > size {
				continue
			}

			remainingSize := size - num

			if dp[row-1][remainingSize] == 1 {
				dp[row][size] = 1
			}
		}
	}

	return dp[len(nums)][capacity] == 1
}

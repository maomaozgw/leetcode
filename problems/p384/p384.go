package p384

import (
	"math/rand"
	"time"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().Unix())
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	var result = make([]int, len(this.nums))
	copy(result, this.nums)
	// rand.Shuffle(len(this.nums), func(i, j int) {
	// 	result[i], result[j] = result[j], result[i]
	// })
	for i := 0; i < len(this.nums); i++ {
		index := rand.Intn(len(this.nums) - i)
		result[len(this.nums)-i-1], result[index] = result[index], result[len(this.nums)-i-1]
	}
	return result
}

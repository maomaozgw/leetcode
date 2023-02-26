package p303

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	var prefixSum []int = make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

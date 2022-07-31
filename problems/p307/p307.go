package p307

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	return NumArray{data: nums}
}

func (this *NumArray) Update(index int, val int) {
	this.data[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int
	for ; left <= right; left++ {
		sum += this.data[left]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

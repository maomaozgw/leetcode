package p307

import "math"

type NumArray struct {
	data      []int
	chunks    []int
	chunkSize int
	calcSize  int
}

func Constructor(nums []int) NumArray {
	n := NumArray{
		data: nums,
	}
	numLen := len(nums)
	n.chunkSize = int(math.Sqrt(float64(numLen)))
	chunkCount := numLen / n.chunkSize
	if numLen > n.chunkSize*chunkCount {
		chunkCount++
	}
	n.chunks = make([]int, chunkCount)
	n.calcSize = n.chunkSize * 2
	for nI := 0; nI < chunkCount; nI++ {
		for i := nI * n.chunkSize; i < (nI+1)*n.chunkSize; i++ {
			if i >= numLen {
				break
			}
			n.chunks[nI] += nums[i]
		}
	}
	return n
}

func (this *NumArray) Update(index int, val int) {
	oldVal := this.data[index]
	this.data[index] = val
	this.chunks[index/this.chunkSize] += val - oldVal
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int
	if right-left < this.calcSize {
		for ; left <= right; left++ {
			sum += this.data[left]
		}
		return sum
	}
	var (
		cs, cr int
	)
	if left%this.chunkSize == 0 {
		cs = left / this.chunkSize
	} else {
		cs = left/this.chunkSize + 1
	}

	cr = right / this.chunkSize

	for ; left < cs*this.chunkSize; left++ {
		sum += this.data[left]
	}
	for ; cs < cr; cs++ {
		sum += this.chunks[cs]
	}
	for r := cr * this.chunkSize; r <= right; r++ {
		sum += this.data[r]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

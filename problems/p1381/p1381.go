package p1381

type CustomStack struct {
	maxSize int
	data    []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		maxSize: maxSize,
		data:    make([]int, 0),
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.data) < this.maxSize {
		this.data = append(this.data, x)
	}

}

func (this *CustomStack) Pop() int {
	if len(this.data) == 0 {
		return -1
	}
	result := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return result
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(this.data); i++ {
		this.data[i] += val
	}
}

package p307

type NumArray struct {
	Data []*Seg
}

type Seg struct {
	Start  int
	Middle int
	End    int
	Data   int
	Mark   int
}

func Constructor(nums []int) NumArray {
	n := NumArray{
		Data: make([]*Seg, len(nums)<<2),
	}
	n.build(0, nums)
	return n
}

func (this *NumArray) build(index int, values []int) {
	node := this.Data[index]
	if node == nil {
		// first node
		node = &Seg{
			Start: 0,
			End:   len(values) - 1,
		}
		this.Data[0] = node
	}
	if node.Start == node.End {
		// 叶子节点
		node.Data = values[node.Start]
		node.Middle = node.Start
		return
	}
	mid := node.Start + (node.End-node.Start)>>1
	node.Middle = mid
	this.Data[(index<<1)+1] = &Seg{
		Start: node.Start,
		End:   mid,
	}
	this.Data[(index<<1)+2] = &Seg{
		Start: mid + 1,
		End:   node.End,
	}
	this.build((index<<1)+1, values)
	this.build((index<<1)+2, values)
	node.Data = this.Data[(index<<1)+1].Data + this.Data[(index<<1)+2].Data
}

func (this *NumArray) Update(index int, val int) {
	this.update(0, index, val)
}

func (this *NumArray) update(nodeIndex, index, value int) {
	node := this.Data[nodeIndex]
	if node == nil {
		return
	}
	if index < node.Start || index > node.End {
		return
	}
	if node.Start == node.End && node.Start == index {
		node.Data = value
		return
	}
	if index <= node.Middle {
		this.update((nodeIndex<<1)+1, index, value)
	} else {
		this.update((nodeIndex<<1)+2, index, value)
	}
	node.Data = this.Data[(nodeIndex<<1)+1].Data + this.Data[(nodeIndex<<1)+2].Data
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(0, left, right)
}

func (this *NumArray) query(index, start, end int) int {
	node := this.Data[index]
	if start <= node.Start && node.End <= end {
		return node.Data
	}
	var result int
	if start <= node.Middle {
		result += this.query(index<<1+1, start, end)
	}
	if end > node.Middle {
		result += this.query(index<<1+2, start, end)
	}
	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

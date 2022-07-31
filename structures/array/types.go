package array

import (
	"golang.org/x/exp/constraints"
)

type Segment[T constraints.Integer | constraints.Float] struct {
	Start  int
	Middle int
	End    int
	Data   T
	Mark   int
}

type SegmentTree[T constraints.Integer | constraints.Float] struct {
	Data    []*Segment[T]
	AggFunc func(tLeft, tRight T) T `json:"-"`
}

func NewSegmentTree[T constraints.Integer | constraints.Float](aggFunc func(tLeft, tRight T) T, values ...T) *SegmentTree[T] {
	var tree = &SegmentTree[T]{
		Data:    nil,
		AggFunc: aggFunc,
	}
	tree.Data = make([]*Segment[T], len(values)<<2)
	tree.build(0, values)
	return tree
}

func (t *SegmentTree[T]) build(index int, data []T) {
	node := t.Data[index]
	if node == nil {
		node = &Segment[T]{
			Start: 0,
			End:   len(data) - 1,
		}
		t.Data[0] = node
	} else if node.Start == node.End {
		// 叶子节点
		node.Data = data[node.Start]
		node.Middle = node.Start
		return
	}
	mid := node.Start + (node.End-node.Start)>>1
	node.Middle = mid
	t.Data[(index<<1)+1] = &Segment[T]{
		Start: node.Start,
		End:   mid,
	}
	t.Data[(index<<1)+2] = &Segment[T]{
		Start: mid + 1,
		End:   node.End,
	}
	t.build((index<<1)+1, data)
	t.build((index<<1)+2, data)
	node.Data = t.AggFunc(t.Data[(index<<1)+1].Data, t.Data[(index<<1)+2].Data)
}

func (t *SegmentTree[T]) Update(index int, value T) {
	t.update(0, index, value)
}
func (t *SegmentTree[T]) update(nodeIndex, index int, value T) {
	node := t.Data[nodeIndex]
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
		t.update((nodeIndex<<1)+1, index, value)
	} else {
		t.update((nodeIndex<<1)+2, index, value)
	}
	node.Data = t.AggFunc(t.Data[(nodeIndex<<1)+1].Data, t.Data[(nodeIndex<<1)+2].Data)
}

func (t *SegmentTree[T]) QueryRange(start, end int) T {
	return t.query(0, start, end)
}

func (t *SegmentTree[T]) query(index, start, end int) T {
	node := t.Data[index]
	if start <= node.Start && node.End <= end {
		return node.Data
	}
	var result T
	if start <= node.Middle {
		result += t.query(index<<1+1, start, end)
	}
	if end > node.Middle {
		result += t.query(index<<1+2, start, end)
	}
	return result
}

package p1797

import (
	"container/heap"
)

type Item struct {
	index  int
	expire int
}

type Heap struct {
	data []*Item
}

// Len is the number of elements in the collection.
func (h *Heap) Len() int {
	return len(h.data)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h *Heap) Less(i int, j int) bool {
	return h.data[i].expire < h.data[j].expire
}

// Swap swaps the elements with indexes i and j.
func (h *Heap) Swap(i int, j int) {
	h.data[i].index, h.data[j].index = j, i
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) Push(x any) {
	item := x.(*Item)
	item.index = len(h.data)
	h.data = append(h.data, item)
}

func (h *Heap) Pop() any {
	n := len(h.data)
	item := h.data[n-1]
	h.data = h.data[:n-1]
	item.index = -1
	return item
}

type AuthenticationManager struct {
	items map[string]*Item
	h     *Heap
	ttl   int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		ttl:   timeToLive,
		items: map[string]*Item{},
		h: &Heap{
			data: []*Item{},
		},
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	item := &Item{expire: this.ttl + currentTime}
	this.items[tokenId] = item
	heap.Push(this.h, item)
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	item, ok := this.items[tokenId]
	if !ok || item.expire <= currentTime {
		return
	}
	item.expire = currentTime + this.ttl
	heap.Fix(this.h, item.index)
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	for this.h.Len() > 0 && this.h.data[0].expire <= currentTime {
		heap.Pop(this.h)
	}
	return this.h.Len()
}

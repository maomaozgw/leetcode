package p460

import "container/heap"

type QueueItem struct {
	priority uint64
	key      uint32
	value    uint32
	index    uint16
}

type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = uint16(i)
	pq[j].index = uint16(j)
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.index = uint16(n)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type LFUCache struct {
	items    []QueueItem
	h        map[uint32]*QueueItem
	queue    PriorityQueue
	priority uint32
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		items: make([]QueueItem, 0, capacity),
		h:     make(map[uint32]*QueueItem, capacity),
		queue: make(PriorityQueue, 0, capacity),
	}
}

func (l *LFUCache) Get(key int) int {
	item := l.h[uint32(key)]
	if item == nil {
		return -1
	}
	item.priority = (((item.priority >> 32) + 1) << 32) | uint64(l.priority)
	l.priority++
	heap.Fix(&l.queue, int(item.index))
	return int(item.value)
}

func (l *LFUCache) Put(key int, value int) {
	if cap(l.items) == 0 {
		return
	}
	k := uint32(key)
	item := l.h[k]
	if item == nil {
		if len(l.items) == cap(l.items) {
			item = l.queue[0]
			item.priority = 0
			delete(l.h, item.key)
		} else {
			n := len(l.items)
			l.items = l.items[:n+1]
			item = &l.items[n]
			l.queue.Push(item)
		}
		item.key = k
		l.h[k] = item
	}
	item.value = uint32(value)
	item.priority = (((item.priority >> 32) + 1) << 32) | uint64(l.priority)
	l.priority++
	heap.Fix(&l.queue, int(item.index))
}

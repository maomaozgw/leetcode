// https://leetcode.com/problems/the-skyline-problem/

package p218

import (
	"container/heap"
	"sort"
)

type edge struct {
	pos    int
	height int
}

type EdgeQueue []*edge

func (q *EdgeQueue) Len() int { return len(*q) }

func (q *EdgeQueue) Less(i, j int) bool { return (*q)[i].height > (*q)[j].height }

func (q *EdgeQueue) Swap(i, j int) { (*q)[i], (*q)[j] = (*q)[j], (*q)[i] }

func (q *EdgeQueue) Push(x interface{}) {
	e := x.(*edge)
	*q = append(*q, e)
}

func (q *EdgeQueue) Pop() interface{} {
	old := *q
	n := len(old)
	e := old[n-1]
	old[n-1] = nil
	*q = old[:n-1]
	return e
}

func getSkyline(buildings [][]int) [][]int {
	var l = len(buildings)
	var positions = make([][]int, l*2)
	// 这个步骤是把原来的矩形变成了点和对应的元素的映射关系
	// [ [] [[] []]  ] -> | || || |  |||    |
	for i := 0; i < l; i++ {
		positions[i*2] = []int{buildings[i][0], i}
		positions[i*2+1] = []int{buildings[i][1], i}
	}
	// 根据点的 X 轴坐标位置排序
	sort.Slice(positions, func(i, j int) bool {
		return positions[i][0] < positions[j][0]
	})

	q := make(EdgeQueue, 0, 2*l)
	heap.Init(&q)
	var idx = 0
	var answer [][]int
	/*
		根据点位置，找到
	*/
	for idx < len(positions) {
		currentIdx := positions[idx][0]
		// 合并与当前点 x 坐标重合的其他点的位置到最大堆中，堆顶端是高度最高的点
		for idx < len(positions) && positions[idx][0] == currentIdx {
			bIdx := positions[idx][1]
			if currentIdx == buildings[bIdx][0] {
				heap.Push(&q, &edge{
					pos:    buildings[bIdx][1],
					height: buildings[bIdx][2],
				})
			}
			idx++
		}
		// 如果最高点的x 坐标已经不再当前点的x 坐标内了，就 POP 掉，同时清理掉其他堆定不在内部的元素
		for len(q) > 0 && q[0].pos <= currentIdx {
			heap.Pop(&q)
		}
		// 最高高度如果堆为空就是0不然就是堆顶的高度
		maxHeight := 0
		if len(q) > 0 {
			maxHeight = q[0].height
		}
		// 如果结果数组为空就添加一个
		// 如果新的最高高度与现有的不一样，也添加
		if len(answer) == 0 || answer[len(answer)-1][1] != maxHeight {
			answer = append(answer, []int{currentIdx, maxHeight})
		}
	}
	return answer

}

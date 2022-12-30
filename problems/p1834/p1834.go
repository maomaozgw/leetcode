// https://leetcode.com/problems/single-threaded-cpu/

package p1834

import "container/heap"

type Heap struct {
	data     []*Task
	lessFunc func(i, j *Task) bool
}

func (h *Heap) Len() int { return len(h.data) }

func (h *Heap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *Heap) Push(x interface{}) { h.data = append(h.data, x.(*Task)) }

func (h *Heap) Peek() *Task {
	return h.data[0]
}

func (h *Heap) Less(i, j int) bool {
	return h.lessFunc(h.data[i], h.data[j])
}

func (h *Heap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

type Task struct {
	Idx            int
	EnqueueTime    int
	ProcessingTime int
}

func getOrder(tasks [][]int) []int {
	// 基于任务的等待时间建立最小堆
	var incomingPQ = &Heap{
		data: []*Task{},
		lessFunc: func(i, j *Task) bool {
			return i.EnqueueTime < j.EnqueueTime
		},
	}
	// 基于任务的执行时间&idx 建立最小堆
	var bufferedTaskPQ = &Heap{
		data: []*Task{},
		lessFunc: func(i, j *Task) bool {
			if i.ProcessingTime == j.ProcessingTime {
				return i.Idx < j.Idx
			}
			return i.ProcessingTime < j.ProcessingTime
		},
	}
	for i := range tasks {
		heap.Push(incomingPQ, &Task{
			Idx:            i,
			EnqueueTime:    tasks[i][0],
			ProcessingTime: tasks[i][1],
		})
	}
	var result []int
	var ticks int

	for incomingPQ.Len() > 0 {
		// 如果当前 CPU 的时间周期里有排队时间满足的任务，就取出所有满足的任务加入到预备执行队列
		for incomingPQ.Len() > 0 && incomingPQ.Peek().EnqueueTime <= ticks {
			heap.Push(bufferedTaskPQ, heap.Pop(incomingPQ))
		}
		// 如果没有可执行的任务，快进 CPU 时间周期到最近的可执行任务的等待时间
		if incomingPQ.Len() > 0 && bufferedTaskPQ.Len() == 0 {
			ticks = incomingPQ.Peek().EnqueueTime
			continue
		}
		processingTask := heap.Pop(bufferedTaskPQ).(*Task)
		ticks += processingTask.ProcessingTime
		result = append(result, processingTask.Idx)
	}
	for bufferedTaskPQ.Len() > 0 {
		processingTask := heap.Pop(bufferedTaskPQ).(*Task)
		result = append(result, processingTask.Idx)
	}
	return result
}

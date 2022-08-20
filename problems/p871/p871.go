// https://leetcode.com/problems/minimum-number-of-refueling-stops/

package p871

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	fuelCnt := 0
	currentFuel := startFuel
	prevPos := 0
	stations = append(stations, []int{target, 100})
	heap := binaryheap.NewWith(func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	})

	for i := 0; i < len(stations); i++ {
		currentFuel -= stations[i][0] - prevPos
		for currentFuel < 0 {
			maxFuel, ok := heap.Pop()
			if !ok {
				return -1
			}
			fuelCnt++
			currentFuel += maxFuel.(int)
		}
		heap.Push(stations[i][1])
		prevPos = stations[i][0]
	}

	return fuelCnt
}

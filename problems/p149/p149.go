package p149

import (
	"math"
)

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return 2
	}

	result := 0
	for i := 0; i < len(points); i++ {
		temp := 0
		res := make(map[float64]int)
		for j := i + 1; j < len(points); j++ {
			r := radio(points[i][0], points[i][1], points[j][0], points[j][1])
			res[r]++
			if temp < res[r] {
				temp = res[r]
			}
		}
		result = max(result, temp)
		if result+1 == len(points)-i {
			return result + 1
		}
	}
	return result + 1
}

func radio(x1, y1, x2, y2 int) float64 {
	return absInf(float64(y2-y1) / float64(x2-x1))
}

func absInf(i float64) float64 {
	if math.IsInf(i, 0) {
		return math.Inf(1)
	}
	return i
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

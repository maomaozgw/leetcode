package p539

import "sort"

func findMinDifference(timePoints []string) int {
	var minuites = make([]int, len(timePoints))
	for i := range timePoints {
		minuites[i] = toMinuites(timePoints[i])
	}
	sort.Ints(minuites)
	var result = 1440
	for i := 1; i < len(minuites); i++ {
		result = min(result, minuites[i]-minuites[i-1])
	}
	result = min(result, minuites[0]+1440-minuites[len(minuites)-1])
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func toMinuites(time string) int {
	return int(time[0]-'0')*600 + int(time[1]-'0')*60 + int(time[3]-'0')*10 + int(time[4]-'0')
}

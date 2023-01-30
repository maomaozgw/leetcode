package p352

type SummaryRanges struct {
	Intervals [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(value int) {
	innerVal := this.Intervals
	len := len(innerVal)
	firstSmaller := func(target, left, right int) int {
		for left < right {
			mid := left + (right-left)/2
			cur := innerVal[mid][0]
			if cur < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
		left--
		return left
	}

	index := firstSmaller(value, 0, len)
	if index+1 < len && innerVal[index+1][0] == value {
		return
	}
	var change [][]int
	var newInterval [][]int
	if index >= 0 {
		change = append(change, innerVal[index])
		newInterval = append(newInterval, innerVal[:index]...)
	}

	change = append(change, []int{value, value})
	if index+1 < len {
		change = append(change, innerVal[index+1])
	}
	change = mergeIntervals(change)

	newInterval = append(newInterval, change...)
	if index+2 < len {
		newInterval = append(newInterval, innerVal[index+2:]...)
	}
	this.Intervals = newInterval
}

func mergeIntervals(input [][]int) [][]int {
	if len(input) == 0 {
		return input
	}
	var result [][]int
	lastInterval := input[0]
	for i := 1; i < len(input); i++ {
		if input[i][0]-lastInterval[1] > 1 {
			result = append(result, lastInterval)
			lastInterval = input[i]
		} else {
			end := lastInterval[1]
			if end < input[i][1] {
				end = input[i][1]
			}
			lastInterval = []int{lastInterval[0], end}
		}
	}
	result = append(result, lastInterval)
	return result
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.Intervals
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */

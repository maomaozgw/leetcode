package p1578

type costs []int

func (c *costs) Len() int {
	return len(*c)
}

func (c *costs) Less(i, j int) bool {
	return (*c)[i] < (*c)[j]
}

func (c *costs) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *costs) Push(x any) {
	*c = append(*c, x.(int))
}

func (c *costs) Pop() any {
	old := *c
	n := len(old)
	item := old[n-1]
	*c = old[0 : n-1]
	return item
}

func minCost(colors string, neededTime []int) int {
	var total int
	var prev byte = colors[0]
	var prevIdx = 0
	var preCost = neededTime[0]
	var maxCost = neededTime[0]
	for i := 1; i < len(colors); i++ {
		if colors[i] == prev {
			preCost += neededTime[i]
			if neededTime[i] > maxCost {
				maxCost = neededTime[i]
			}
		} else {
			if i-prevIdx > 1 {
				// 有重复气球
				total += preCost - maxCost
			}
			prevIdx = i
			prev = colors[i]
			preCost = neededTime[i]
			maxCost = neededTime[i]
		}
	}

	if len(colors)-prevIdx > 1 {
		total += preCost - maxCost
	}
	return total
}

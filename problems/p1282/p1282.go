package p1282

func groupThePeople(groupSizes []int) [][]int {
	var group = map[int][]int{}
	var result [][]int
	for i, num := range groupSizes {
		if _, ok := group[num]; !ok {
			group[num] = []int{}
		}
		group[num] = append(group[num], i)
		if len(group[num]) == num {
			result = append(result, group[num])
			group[num] = []int{}
		}
	}
	return result
}

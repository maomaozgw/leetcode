package p1128

const (
	offset = 100
)

func numEquivDominoPairs(dominoes [][]int) int {
	var eleMap = map[int]int{}

	for _, ele := range dominoes {
		eleMap[key(ele)]++
	}
	var result int
	for _, v := range eleMap {
		result += v * (v - 1) / 2
	}
	return result
}

func key(items []int) int {
	if items[0] > items[1] {
		return items[0]*offset + items[1]
	} else {
		return items[1]*offset + items[0]
	}
}

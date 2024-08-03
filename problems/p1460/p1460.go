package p1460

func canBeEqual(target []int, arr []int) bool {
	var cntMap = make(map[int]int)
	for _, num := range target {
		cntMap[num]++
	}
	for _, num := range arr {
		if cntMap[num] == 0 {
			return false
		}
		cntMap[num]--
	}
	return true
}

package p904

func totalFruit(fruits []int) int {
	var fMap = map[int]int{}
	var i, result = 0, 0
	for j := 0; j < len(fruits); j++ {
		fMap[fruits[j]]++
		if len(fMap) <= 2 {
			newRes := j - i + 1
			if newRes > result {
				result = newRes
			}
		} else {
			fMap[fruits[i]]--
			if fMap[fruits[i]] == 0 {
				delete(fMap, fruits[i])
			}
			i++
		}
	}
	return result
}

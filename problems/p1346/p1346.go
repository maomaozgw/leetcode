package p1346

func checkIfExist(arr []int) bool {
	var m = map[int]bool{}
	var zCnt = 0
	for _, num := range arr {

		m[num] = true
		if num == 0 {
			zCnt++
			if zCnt > 1 {
				return true
			}
			continue
		}

		if m[num*2] || (num%2 == 0 && m[num/2]) {
			return true
		}

	}
	return false
}

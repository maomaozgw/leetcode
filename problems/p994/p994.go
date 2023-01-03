package p994

func minDeletionSize(strs []string) int {
	var l = len(strs)

	if l <= 1 {
		return 0
	}
	var n = len(strs[0])
	var deleteCnt = 0
	var deletedCache = make([]bool, n)
	for i := 1; i < l; i++ {
		for j := 0; j < n; j++ {
			if deletedCache[j] || strs[i][j] >= strs[i-1][j] {
				continue
			}
			deleteCnt++
			deletedCache[j] = true
			if deleteCnt == n {
				return n
			}
		}
	}
	return deleteCnt
}

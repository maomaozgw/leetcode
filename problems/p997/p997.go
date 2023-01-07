package p997

func findJudge(n int, trust [][]int) int {
	var trustedCnt = make([]int, n+1)
	var trustCnt = make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		trustCnt[trust[i][0]]++
		trustedCnt[trust[i][1]]++
	}
	var judgeNum = -1
	var judgeCnt = 0
	for i := 1; i <= n; i++ {
		if trustCnt[i] != 0 {
			continue
		}
		judgeNum = i
		judgeCnt++
	}
	if judgeCnt != 1 {
		return -1
	}
	if trustedCnt[judgeNum] != n-1 {
		return -1
	}
	return judgeNum
}

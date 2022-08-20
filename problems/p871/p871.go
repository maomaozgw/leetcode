// https://leetcode.com/problems/minimum-number-of-refueling-stops/

package p871

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	l := len(stations)
	var dp = make([]int, l+1)
	dp[0] = startFuel
	for i := 0; i < l; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] < stations[i][0] {
				continue
			}
			tmp := dp[j] + stations[i][1]
			if dp[j+1] < tmp {
				dp[j+1] = tmp
			}
		}
		if stations[i][0] > target {
			break
		}
	}
	for i := 0; i < l+1; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}

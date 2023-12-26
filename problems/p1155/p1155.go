// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/

package p1155

const (
	modVal = 1e9 + 7
)

func numRollsToTarget(n int, k int, target int) int {
	dp := make([]int, target)
	// 第一次 roll ，最多只能到 min(k,target)
	for j := 0; j < min(k, target); j++ {
		dp[j] = 1
	}
	// roll n 次
	// 如果 k = 5 target = 8
	// ROLL/TAR 1   2   3   4   5   6   7   8
	// 0        1   1   1   1   1   0   0   0
	// 1        0   1   2   3   4   4   5   4
	// 2        0   0   1   3   6   10  15  19
	// f(5) = f(4') + f(3') + f(2') + f(1')
	for i := 1; i < n; i++ {
		// 最小值为 1，最差每次都是 1
		var tmp = make([]int, target)
		for j := i; j < target; j++ {
			for l := 1; l <= min(target, k); l++ {
				if j-l < 0 {
					break
				}
				tmp[j] += dp[j-l]
				tmp[j] %= modVal
			}
		}
		dp = tmp
	}
	return dp[target-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

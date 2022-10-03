// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/

package p1155

const (
	modVal = 1e9 + 7
)

func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n)
	// 记录每次 roll 的结果数据
	for i := 0; i < n; i++ {
		// 每次 roll 到 target + 1 需要的次数
		dp[i] = make([]int, target)
	}
	// 第一次 roll ，最多只能到 min(k,target)
	for j := 0; j < min(k, target); j++ {
		dp[0][j] = 1
	}
	// roll n 次
	// 如果 k = 5 target = 8
	// ROLL/TAR 1   2   3   4   5   6   7   8
	// 0        1   1   1   1   1   0   0   0
	// 1        0   1   2   3   4   4   5   4
	// 2        0   0   1   3   6   10  15  19
	// 如果目标值不大于 k + 1，每 roll 一次总和为目标值需要的次数等于
	// f(5) = f(4') + f(3') + f(2') + f(1')
	// f(4) =  f(3') + f(2') + f(1')
	// f(5) = f(4) + f(4')
	// f(6) = f(5) + f(5')
	// 如果目标值大于 k + 1 ， roll 一次总和为目标值需要的次数等于
	// f(7) = f(6') + f(5') + f(4') + f(3') + f(2')
	// f(7) = f(6) + f(6') - f(1')
	// f(8) = f(7') + f(6') + f(5') + f(4') + f(3')
	// f(8) = f(7) + f(7') - f(2')
	// 汇总后公式如下
	// f(n) = f(n-1) + f((n-1)')
	// f(n) = f(n-1) + f((n-1)') - f((n-(k+1))')
	for i := 1; i < n; i++ {
		// 最小值为 1，最差每次都是 1
		for j := i; j < target; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			if j > k {
				dp[i][j] -= dp[i-1][j-k-1]
			}

			dp[i][j] %= modVal
		}
	}
	return dp[n-1][target-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

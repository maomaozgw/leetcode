// https://leetcode.com/problems/decode-ways/

package p91

func numDecodings(s string) int {
	var dp = make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			dp[i+1] += dp[i]
		}
		if i > 0 && (s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6') {
			dp[i+1] += dp[i-1]
		}

	}
	return dp[len(s)]
}

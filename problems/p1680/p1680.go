// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/

package p1680

import (
	"math"
)

const (
	maxVal = 1000000007
)

var (
	result = make([]int, 100001)
)

func init() {
	for i := int64(1); i < 100001; i++ {
		n := int(math.Log2(float64(i))) + 1
		result[i] = int(i + int64(result[i-1])<<n%int64(maxVal))
	}
}

func concatenatedBinary(n int) int {
	return result[n]
}

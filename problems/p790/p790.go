package p790

/*
0 -> 1
1 -> 1
2 -> 2
3 -> 5      2*2 + 1
4 -> 11     5*2 + 1
5 -> 24     11*2 + 2
6 -> 53     24*2 + 5
7 -> 117    53*2 + 11
if (i == 0) return 1;
if (i == 1) return 1;
if (i == 2) return 2;
return f(i - 1) * 2 + f(i - 3);
*/

const (
	modVal = 1000000007
	maxVal = 1000
)

var (
	cache []int
)

func init() {
	cache = make([]int, maxVal+1)
	cache[0] = 1
	cache[1] = 1
	cache[2] = 2
	for i := 3; i <= maxVal; i++ {
		cache[i] = (cache[i-1]*2 + cache[i-3]) % modVal
	}
}
func numTilings(n int) int {
	return cache[n]
}

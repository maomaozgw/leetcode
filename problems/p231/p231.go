package p231

var power2Map = map[int]bool{}

func init() {
	for i := 0; i < 32; i++ {
		power2Map[1<<i] = true
	}
}

func isPowerOfTwo(n int) bool {
	return power2Map[n]
}

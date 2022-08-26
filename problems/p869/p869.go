package p869

import (
	"sort"
	"strconv"
)

var table map[string]bool

func init() {
	table = map[string]bool{}
	for i := 1; i <= 1000000000; i <<= 1 {
		table[intKey(i)] = true
	}
}

func intKey(n int) string {
	bytes := []byte(strconv.Itoa(n))
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func reorderedPowerOf2(n int) bool {
	return table[intKey(n)]
}

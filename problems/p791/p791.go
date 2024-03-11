package p791

import (
	"sort"
)

func customSortString(order string, s string) string {
	var ref = make([]int, 26)
	for i := range order {
		ref[order[i]-'a'] = i + 1
	}
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		v1, v2 := ref[b[i]-'a'], ref[b[j]-'a']
		if v1 == v2 {
			return b[i] < b[j]
		}
		return v1 < v2
	})
	return string(b)
}

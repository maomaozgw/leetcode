// https://leetcode.com/problems/minimum-genetic-mutation/

package p433

var (
	gens = []byte{'A', 'C', 'G', 'T'}
)

func minMutation(start string, end string, bank []string) int {
	var level = 0
	var bankMap = map[string]bool{}
	for _, b := range bank {
		bankMap[b] = true
	}
	var q []string
	q = append(q, start)
	for len(q) > 0 {
		level++
		var newQ []string
		for i := 0; i < len(q); i++ {
			origin := []byte(q[i])
			for j := 0; j < 8; j++ {
				var originCc = origin[j]
				for k := 0; k < 4; k++ {
					origin[j] = gens[k]
					originS := string(origin)
					if _, ok := bankMap[originS]; !ok {
						continue
					}
					if originS == end {
						return level
					}
					delete(bankMap, originS)
					newQ = append(newQ, originS)

				}
				origin[j] = originCc
			}
		}
		q = newQ
	}
	return -1
}

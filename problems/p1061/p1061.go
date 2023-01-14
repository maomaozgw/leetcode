package p1061

const (
	start = 'a'
)

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	var mapping = make([]byte, 26)
	for i := 0; i < 26; i++ {
		mapping[i] = byte(i)
	}
	var find = func(i byte) (root byte) {
		for mapping[i] != i {
			i = mapping[i]
		}
		return i
	}
	var union = func(i, j byte) {
		r1, r2 := find(i), find(j)
		if r1 != r2 {
			if r1 > r2 {
				mapping[r1] = r2
			} else {
				mapping[r2] = r1
			}
		}
	}
	var l = len(s1)
	if l == 0 {
		return baseStr
	}
	for i := 0; i < l; i++ {
		union(s1[i]-start, s2[i]-start)
	}
	var sb = []byte(baseStr)
	for i := 0; i < len(sb); i++ {
		sb[i] = find(sb[i]-start) + start
	}
	return string(sb)

}

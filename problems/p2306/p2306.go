package p2306

const (
	start = byte('a')
	end   = byte('z')
)

func distinctNames(ideas []string) int64 {
	var prefixMap = map[byte]map[string]bool{}

	for idx := range ideas {
		prefix := ideas[idx][0]
		suffix := ideas[idx][1:]
		if _, exists := prefixMap[prefix]; !exists {
			prefixMap[prefix] = map[string]bool{}
		}
		prefixMap[prefix][suffix] = true
	}
	var result int64 = 0

	for i := start; i < end; i++ {
		startSuffixes, exists := prefixMap[i]
		if !exists {
			continue
		}
		for j := i + 1; j <= end; j++ {
			endSuffixes, exists := prefixMap[j]
			if !exists {
				continue
			}
			iCnt, jCnt := 0, 0
			for key := range startSuffixes {
				if !endSuffixes[key] {
					iCnt++
				}
			}
			for key := range endSuffixes {
				if !startSuffixes[key] {
					jCnt++
				}
			}
			result += int64(iCnt) * int64(jCnt) * 2
		}
	}
	return result
}

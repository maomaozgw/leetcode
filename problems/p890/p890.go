// https://leetcode.com/problems/find-and-replace-pattern/

package p890

type Pattern struct {
	firstIdx int
	index    []int
}

func findAndReplacePattern(words []string, pattern string) []string {
	var result []string

	for _, w := range words {
		if len(pattern) != len(w) {
			continue
		}
		mapping := map[uint8]uint8{}
		backMapping := map[uint8]uint8{}
		match := true
		for idx := 0; idx < len(w); idx++ {
			r := w[idx]
			rr := pattern[idx]
			if mr, ok := mapping[r]; !ok {
				if _, ok := backMapping[rr]; ok {
					match = false
					break
				}
				mapping[r] = rr
				backMapping[rr] = r
			} else if mr != rr {
				match = false
				break
			}
		}
		if !match {
			continue
		}
		result = append(result, w)
	}
	return result
}

func findAndReplacePattern2(words []string, pattern string) []string {
	var result []string
	patternLen := len(pattern)
	if len(pattern) <= 1 {
		for _, w := range words {
			if len(w) == patternLen {
				result = append(result, w)
			}
		}
		return result
	}

	patternSlice := extractPattern(pattern)

	for _, w := range words {
		if len(w) != patternLen {
			continue
		}
		var match = true
		wSlice := extractPattern(w)
		if len(patternSlice) != len(wSlice) {
			continue
		}
		for idx, p := range patternSlice {
			wp := wSlice[idx]
			if wp.firstIdx != p.firstIdx {
				match = false
				break
			}
			if len(wp.index) != len(p.index) {
				match = false
				break
			}
			for ii, v := range p.index {
				if wp.index[ii] != v {
					match = false
					break
				}
			}
			if !match {
				break
			}
		}
		if !match {
			continue
		}
		result = append(result, w)
	}
	return result
}

func extractPattern(w string) []*Pattern {
	var patternMap = map[rune]*Pattern{}
	var patternSlice []*Pattern
	for wIdx, r := range w {
		if info, ok := patternMap[r]; !ok {
			p := &Pattern{
				firstIdx: wIdx,
				index:    []int{},
			}
			patternMap[r] = p
			patternSlice = append(patternSlice, p)
		} else {
			info.index = append(info.index, wIdx)
		}
	}

	return patternSlice
}

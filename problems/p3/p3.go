package p3

func lengthOfLongestSubstring2(s string) int {
	length := 0
	longest := 0
	cacheMap := map[uint8]int{}
	startIndex := 0
	for index := 0; index < len(s); index++ {
		c := s[index]
		length = index - startIndex + 1
		if prevIndex, exists := cacheMap[c]; exists {
			cacheMap[c] = index
			if prevIndex >= startIndex {
				length -= 1
				startIndex = prevIndex + 1
			}
			if length > longest {
				longest = length
			}
		} else {
			cacheMap[c] = index
		}
	}
	if length > longest {
		longest = length
	}
	return longest
}

func lengthOfLongestSubstring(s string) int {
	var result = 0
	var currentCnt = 0
	var currentCharSet = make([]int, 128)
	for i := 0; i < 128; i++ {
		currentCharSet[i] = -1
	}
	var lastIdx = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		currentCnt = i - lastIdx + 1
		if currentCharSet[c] > -1 {
			if currentCharSet[c] >= lastIdx {
				currentCnt--
				lastIdx = currentCharSet[c] + 1
			}
			result = max(result, currentCnt)
		}
		currentCharSet[c] = i
	}
	result = max(result, currentCnt)
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

package p443

func compress(chars []byte) int {
	var l = len(chars)
	if l <= 1 {
		return l
	}
	var left = 0
	var prev = chars[0]
	var prevCnt = 1
	for i := 1; i < l; i++ {
		if chars[i] == prev {
			prevCnt++
			continue
		}
		chars[left] = prev
		left++
		prev = chars[i]
		if prevCnt > 1 {
			left = writeCnt(chars, left, prevCnt)
		}
		prevCnt = 1
	}
	chars[left] = prev
	left++
	if prevCnt > 1 {
		left = writeCnt(chars, left, prevCnt)
	}
	return left
}

func writeCnt(chars []byte, start, cnt int) int {
	var left = start
	for cnt > 0 {
		chars[left] = byte(cnt%10) + '0'
		cnt /= 10
		left++
	}
	reverse(chars, start, left-1)
	return left
}

func reverse(chars []byte, i, j int) {
	for j-i >= 1 {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}
}

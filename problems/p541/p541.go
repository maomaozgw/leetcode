package p541

func reverseStr(s string, k int) string {
	content := []byte(s)
	var i, l = 0, len(content)
	var k2 = k * 2
	var cnt = l / k2
	for j := 0; j < cnt; j++ {
		reverse(content, i, i+k-1)
		i += k2
	}
	var remain = l % k2
	if remain == 0 {
		return string(content)
	}

	if remain >= k {
		reverse(content, i, i+k-1)
	} else {
		reverse(content, i, l-1)
	}

	return string(content)
}

func reverse(content []byte, start, end int) {
	if start == end {
		return
	}
	for {
		content[start], content[end] = content[end], content[start]
		if end-start <= 1 {
			return
		}
		start++
		end--
	}
}

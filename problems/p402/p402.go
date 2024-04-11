package p402

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	stack := make([]byte, 0, len(num))
	for i := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	stack = stack[:len(stack)-k]

	i := 0
	for i < len(stack) && stack[i] == '0' {
		i++
	}

	if i == len(stack) {
		return "0"
	}

	return string(stack[i:])
}

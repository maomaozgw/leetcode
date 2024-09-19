package p241

import "strconv"

func diffWaysToCompute(expression string) []int {
	var result []int
	for i := range expression {
		op := expression[i]
		if op == '+' || op == '-' || op == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch op {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}
	if len(result) == 0 {
		val, _ := strconv.Atoi(expression)
		result = append(result, val)
	}
	return result
}

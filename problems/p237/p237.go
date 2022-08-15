package p237

import "strings"

const (
	hundred = "Hundred"
	zero    = "Zero"
)

var (
	tenSlice     = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	hundredSlice = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	bigSlice     = []string{"", "Thousand", "Million", "Billion", "Trillion"}
)

func numberToWords(num int) string {
	if num == 0 {
		return zero
	}
	bigIndex := 0
	var result string
	for num > 0 {
		one := numberToHundred(num % 1000)
		if len(one) > 0 {
			one += " " + bigSlice[bigIndex]
			result = one + " " + result
		}
		num /= 1000
		bigIndex++
	}
	return strings.TrimSpace(result)
}

func numberToHundred(num int) string {
	var result strings.Builder
	h := num / 100
	if h > 0 {
		result.WriteString(tenSlice[h])
		result.WriteString(" ")
		result.WriteString(hundred)
		result.WriteString(" ")
	}
	t := num % 100
	if t > 19 {
		result.WriteString(hundredSlice[t/10])
		result.WriteString(" ")
		result.WriteString(tenSlice[t%10])
	} else {
		result.WriteString(tenSlice[t])
	}
	return strings.TrimSpace(result.String())
}

package p43

import (
	"bytes"
)

const (
	zero = byte('0')
)

var (
	zStr = string(zero)
)

func multiply(num1, num2 string) string {
	// `num1[i] * num2[j]` will be placed at indices `[i + j`, `i + j + 1]`
	if num1 == "0" || num2 == "0" {
		return zStr
	}
	var l1, l2 = len(num1), len(num2)
	var result = bytes.Repeat([]byte{zero}, l1+l2)
	for i := l1 - 1; i > -1; i-- {
		for j := l2 - 1; j > -1; j-- {
			p1 := i + j
			p2 := i + j + 1
			sum := (num1[i]-zero)*(num2[j]-zero) + result[p2] - zero
			result[p1] += sum / 10
			result[p2] = sum%10 + zero
		}
	}
	result = bytes.TrimLeft(result, "0")
	if len(result) == 0 {
		return zStr
	}
	return string(result)
}

func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return string(zero)
	}
	l1 := len(num1)
	l2 := len(num2)
	var l int
	var num = num1
	var bs []byte
	if l1 < l2 {
		l = l1
		bs = []byte(num2)
		num = num1
	} else {
		l = l2
		bs = []byte(num1)
		num = num2
	}
	var result []byte
	for i := l - 1; i > -1; i-- {
		c := num[i]
		result = addStrings(result, multiplySingle(bs, c, l-i-1))
	}
	if len(result) == 0 {
		return string(zero)
	}
	return string(result)
}

func multiplySingle(num1 []byte, c byte, padding int) []byte {
	if c == zero {
		return []byte{}
	}
	var flag = byte(0)
	var result = make([]byte, len(num1)+1+padding)
	for i := len(num1); i < len(result); i++ {
		result[i] = zero
	}
	for i := len(num1) - 1; i > -1; i-- {
		tmp := (num1[i]-zero)*(c-zero) + flag
		flag = tmp / 10
		tmp %= 10
		result[i+1] = tmp + zero
	}
	result[0] = flag + zero
	if result[0] == zero {
		return result[1:]
	}
	return result
}

func addStrings(num1 []byte, num2 []byte) []byte {
	var n1 = len(num1)
	var n2 = len(num2)
	var result []byte
	if n1 > n2 {
		result = make([]byte, n1+1)
	} else {
		result = make([]byte, n2+1)
	}
	var l = len(result)
	var i = 1
	var flag = byte(0)
	for ; i <= n1 && i <= n2; i++ {
		tmp := num1[n1-i] - zero + num2[n2-i] - zero + flag
		if tmp > 9 {
			tmp %= 10
			flag = 1
		} else {
			flag = 0
		}
		result[l-i] = tmp + zero
	}

	for ; i <= n1; i++ {
		tmp := num1[n1-i] - zero + flag
		if tmp > 9 {
			tmp %= 10
			flag = 1
		} else {
			flag = 0
		}
		result[l-i] = tmp + zero
	}

	for ; i <= n2; i++ {
		tmp := num2[n2-i] - zero + flag
		if tmp > 9 {
			tmp %= 10
			flag = 1
		} else {
			flag = 0
		}
		result[l-i] = tmp + zero
	}
	result[0] = flag + zero
	if result[0] == zero {
		return result[1:]
	}
	return result
}

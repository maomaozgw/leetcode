package tools

import (
	"strconv"
	"strings"
)

func Ptr[T comparable](val T) *T {
	tmp := val
	return &tmp
}

func NewGridFromStr(s string) [][]int {
	const (
		sep = ",["
	)
	ms := strings.Split(s, sep)
	var result = make([][]int, len(ms))
	for idx, m := range ms {
		ns := strings.Split(m, ",")
		result[idx] = make([]int, len(ns))
		for ndx, n := range ns {
			n = strings.Trim(n, "[]")
			nv, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			result[idx][ndx] = nv
		}
	}
	return result
}

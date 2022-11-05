package tools

import (
	"strings"
)

func Ptr[T comparable](val T) *T {
	tmp := val
	return &tmp
}

func NewGridFromStr[T any](convert func(string) (T, error), s string) [][]T {
	const (
		sep = ",["
	)
	ms := strings.Split(s, sep)
	var result = make([][]T, len(ms))
	for idx, m := range ms {
		ns := strings.Split(m, ",")
		result[idx] = make([]T, len(ns))
		for ndx, n := range ns {
			n = strings.Trim(n, "[]")
			nv, err := convert(n)
			if err != nil {
				panic(err)
			}
			result[idx][ndx] = nv
		}
	}
	return result
}

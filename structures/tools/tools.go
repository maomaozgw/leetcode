package tools

import (
	"fmt"
	"strings"
)

func Ptr[T comparable](val T) *T {
	tmp := val
	return &tmp
}

func ByteConverter(s string) (byte, error) {
	s = strings.Trim(s, "\"")
	if len(s) != 1 {
		return 0, fmt.Errorf("invalid input %s", s)
	}
	return s[0], nil

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
			n = strings.Trim(n, "[]\n\t")
			if len(n) == 0 {
				continue
			}
			nv, err := convert(n)
			if err != nil {
				panic(err)
			}
			result[idx][ndx] = nv
		}
	}
	return result
}

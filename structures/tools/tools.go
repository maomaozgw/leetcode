package tools

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/exp/constraints"
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
		sep    = "],"
		closet = "[]"
	)
	ms := strings.Split(strings.ReplaceAll(s, "\n", ""), sep)
	var result = make([][]T, len(ms))
	for idx, m := range ms {
		m = strings.TrimSpace(m)
		if m == closet {
			continue
		}
		m = strings.Trim(m, closet)
		if len(m) == 0 {
			continue
		}
		ns := strings.Split(m, ",")
		result[idx] = make([]T, len(ns))
		for ndx, n := range ns {
			n = strings.TrimSpace(n)
			if n == closet {
				continue
			}
			n = strings.TrimSpace(strings.Trim(n, closet))
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

func GridEqual[T constraints.Ordered](one, another [][]T) bool {
	return cmp.Equal(one, another, cmpopts.EquateEmpty(), cmpopts.SortSlices(func(a, b []T) bool {
		if len(a) != len(b) {
			return len(a) < len(b)
		}
		for i := range a {
			if a[i] == b[i] {
				continue
			} else if a[i] < b[i] {
				return true
			}
			return false
		}
		return false
	}))
}

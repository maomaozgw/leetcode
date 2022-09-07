package treenode

import "strconv"

type Converter[T comparable] func(string) T

var (
	MustInt Converter[int] = func(str string) int {
		val, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return val
	}
)

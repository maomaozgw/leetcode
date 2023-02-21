package p709

import "unsafe"

const (
	start = 'A' - 1
	end   = 'Z' + 1

	offSet = 'a' - 'A'
)

func toLowerCase(s string) string {
	var bs = []byte(s)
	for i := 0; i < len(bs); i++ {
		if bs[i] < end && bs[i] > start {
			bs[i] += offSet
		}
	}
	return *(*string)(unsafe.Pointer(&bs))
}

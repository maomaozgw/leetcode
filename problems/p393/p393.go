// https://leetcode.com/problems/utf-8-validation/

package p393

const (
	One   = 1<<7 + 0
	Multi = 1<<6 + 128
	Two   = 1<<5 + 128 + 64
	Three = 1<<4 + 128 + 64 + 32
	Four  = 1<<3 + 128 + 64 + 32 + 16
)

func validUtf8(data []int) bool {
	multiCount := 0
	for i := 0; i < len(data); i++ {
		val := data[i]
		if multiCount > 0 {
			if val >= One && val <= Multi {
				multiCount--
				continue
			}
			return false
		} else if val < One {
			multiCount = 0
		} else if val < Two {
			multiCount = 1
		} else if val < Three {
			multiCount = 2
		} else if val < Four {
			multiCount = 3
		} else {
			return false
		}
	}
	return multiCount == 0
}

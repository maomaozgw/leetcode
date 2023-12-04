package p2264

const (
	startC = byte('/')
	lastC  = byte('9')
)

func largestGoodInteger(num string) string {
	var prvC = num[0]
	var cnt = 1
	var maxC = startC
	for i := 1; i < len(num); i++ {
		if num[i] == prvC {
			cnt++
			if cnt >= 3 {
				maxC = max(maxC, prvC)
				if maxC == lastC {
					break
				}
			}
			continue
		}
		prvC = num[i]
		cnt = 1
	}
	if maxC == startC {
		return ""
	}
	return string([]byte{maxC, maxC, maxC})
}

func max(i, j byte) byte {
	if i > j {
		return i
	}
	return j
}

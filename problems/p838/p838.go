// https://leetcode.com/problems/push-dominoes/

package p838

func pushDominoes(dominoes string) string {
	var result = []byte(dominoes)
	var pendingStart = -1
	var pendingForce = '.'
	for i := 0; i < len(result); i++ {
		c := result[i]
		result[i] = c
		switch c {
		case '.':
			if pendingStart < 0 {
				pendingStart = i
			}
		case 'L':
			if pendingStart >= 0 {
				if pendingForce == '.' {
					for j := pendingStart; j < i; j++ {
						result[j] = 'L'
					}
				} else {
					for j := 0; j < (i-pendingStart)/2; j++ {
						result[pendingStart+j] = 'R'
						result[i-1-j] = 'L'
					}
				}
				pendingStart = -1
			}
			pendingForce = '.'
		case 'R':
			if pendingStart >= 0 {
				if pendingForce == 'R' {
					for j := pendingStart; j < i; j++ {
						result[j] = 'R'
					}
				}
				pendingStart = -1
			}
			pendingForce = 'R'
		}
	}
	if pendingStart > 0 && pendingForce == 'R' {
		for j := pendingStart; j < len(result); j++ {
			result[j] = 'R'
		}
	}

	return string(result)
}

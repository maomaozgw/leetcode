// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/

package p1239

func maxLength(arr []string) int {
	return findMax(arr, 0, 0)
}

func findMax(arr []string, idx int, roadMap int) int {
	if idx == len(arr) {
		return 0
	}
	var prevL = findMax(arr, idx+1, roadMap)
	for i := 0; i < len(arr[idx]); i++ {
		cv := 1 << (arr[idx][i] - 'a')
		if roadMap&cv == cv {
			return prevL
		}
		roadMap |= cv
	}
	return max(prevL, len(arr[idx])+findMax(arr, idx+1, roadMap))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package p34

func searchRange(nums []int, target int) []int {
	const (
		notFound = -1
	)
	var (
		firstIndex = notFound
		lastIndex  = notFound
	)

	for idx, val := range nums {
		if val > target {
			break
		}
		if val != target {
			continue
		}

		if firstIndex == notFound {
			firstIndex = idx
		}
		lastIndex = idx
	}
	return []int{firstIndex, lastIndex}
}

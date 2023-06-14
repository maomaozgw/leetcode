package p228

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var result []string
	var prev = nums[0]
	var prevStart = prev
	for i := 1; i < len(nums); i++ {
		if nums[i]-prev == 1 {
			prev = nums[i]
			continue
		}
		if prev == prevStart {
			result = append(result, strconv.Itoa(prevStart))
		} else {
			result = append(result, strconv.Itoa(prevStart)+"->"+strconv.Itoa(prev))
		}
		prev = nums[i]
		prevStart = prev
	}
	if prev == prevStart {
		result = append(result, strconv.Itoa(prevStart))
	} else {
		result = append(result, strconv.Itoa(prevStart)+"->"+strconv.Itoa(prev))
	}
	return result
}

package p896

func isMonotonic(nums []int) bool {
	var asc *bool = nil
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if asc == nil {
			result := nums[i] > nums[i-1]
			asc = &result
		} else {
			if (nums[i] > nums[i-1]) != *asc {
				return false
			}
		}
	}
	return true
}

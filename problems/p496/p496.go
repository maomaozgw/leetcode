package p496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var numMap = map[int]int{}
	var rightMaxVal = -1
	var rightMaxIndex = -1
	for i := len(nums2) - 1; i >= 0; i-- {
		if nums2[i] > rightMaxVal {
			rightMaxVal = nums2[i]
			rightMaxIndex = i
			numMap[nums2[i]] = -1
			continue
		}
		for j := i + 1; j <= rightMaxIndex; j++ {
			if nums2[j] > nums2[i] {
				numMap[nums2[i]] = nums2[j]
				break
			}
		}
	}
	var result = make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = numMap[nums1[i]]
	}
	return result
}

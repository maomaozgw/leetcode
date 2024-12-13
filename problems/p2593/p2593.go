package p2593

import "sort"

type item struct {
	val int
	idx int
}

func findScore(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	var result int64
	var items []item = make([]item, len(nums))
	for i := range nums {
		it := item{
			val: nums[i],
			idx: i,
		}
		items[i] = it
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].val == items[j].val {
			return items[i].idx < items[j].idx
		}
		return items[i].val < items[j].val
	})
	for i := range items {
		if nums[items[i].idx] == 0 {
			continue
		}
		result += int64(items[i].val)
		switch items[i].idx {
		case 0:
			nums[items[i].idx+1] = 0
		case len(nums) - 1:
			nums[items[i].idx-1] = 0
		default:
			nums[items[i].idx-1] = 0
			nums[items[i].idx+1] = 0
		}
	}
	return result
}

package p729

type MyCalendar struct {
	Items []item
}

type item struct {
	Start int
	End   int
}

func Constructor() MyCalendar {
	return MyCalendar{
		Items: []item{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.Items) == 0 {
		this.Items = append(this.Items, item{Start: start, End: end})
		return true
	}
	for _, i := range this.Items {
		if i.Compare(start, end) == 0 {
			return false
		}
	}
	this.Items = append(this.Items, item{Start: start, End: end})
	return true
}

func (i item) Compare(start, end int) int {
	if i.Start >= end {
		return -1
	} else if i.End <= start {
		return 1
	}
	return 0
}

func (this *MyCalendar) binaryFind(leftIndex, rightIndex int, start, end int) bool {
	mid := (rightIndex + leftIndex) / 2
	if mid == leftIndex {
		if this.Items[mid].Compare(start, end) == 0 {
			return true
		}
		if this.Items[rightIndex].Compare(start, end) == 0 {
			return true
		}
		return false
	}
	midItem := this.Items[mid]
	result := midItem.Compare(start, end)
	if result == 0 {
		return true
	} else if result < 0 {
		return this.binaryFind(leftIndex, mid, start, end)
	} else {
		return this.binaryFind(mid, rightIndex, start, end)
	}
}

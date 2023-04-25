package p2336

type SmallestInfiniteSet struct {
	valCounter []int
}

func Constructor() SmallestInfiniteSet {
	tmp := make([]int, 10001)
	for i := 1; i <= 1000; i++ {
		tmp[i] = 1
	}
	return SmallestInfiniteSet{
		valCounter: tmp,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := range this.valCounter {
		if this.valCounter[i] > 0 {
			this.valCounter[i]--
			return i
		}
	}
	return -1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	this.valCounter[num] = 1
}

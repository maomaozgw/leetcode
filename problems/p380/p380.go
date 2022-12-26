// https://leetcode.com/problems/insert-delete-getrandom-o1/

package p380

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	rnd  *rand.Rand
	keys []int
	data map[int]int
}

func Constructor() RandomizedSet {
	r := RandomizedSet{
		data: map[int]int{},
		rnd:  rand.New(rand.NewSource(time.Now().Unix())),
	}

	return r
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.data[val]; ok {
		return false
	}
	this.data[val] = len(this.keys)
	this.keys = append(this.keys, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.data[val]; ok {
		lastVal := this.keys[len(this.keys)-1]
		this.keys[idx] = lastVal
		this.keys = this.keys[:len(this.keys)-1]
		this.data[lastVal] = idx
		delete(this.data, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	idx := this.rnd.Intn(len(this.keys))
	return this.keys[idx]
}

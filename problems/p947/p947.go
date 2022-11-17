// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/

package p947

const (
	maxLen = 10001
)

type UnionFinder struct {
	Data []int
}

func NewUnionFinder(l int) *UnionFinder {
	u := &UnionFinder{Data: make([]int, l)}
	for i := range u.Data {
		u.Data[i] = i
	}
	return u
}

func (u *UnionFinder) Find(i int) (root int) {
	for u.Data[i] != i {
		i = u.Data[i]
	}
	return i
}

func (u *UnionFinder) Union(i, j int) {
	r1, r2 := u.Find(i), u.Find(j)
	if r1 != r2 {
		u.Data[r2] = r1
	}
}

func (u *UnionFinder) Count() (c int) {
	var visited [maxLen]bool
	for _, i := range u.Data {
		r := u.Find(i)
		if visited[r] {
			continue
		}
		c++
		visited[r] = true
	}
	return
}

func removeStones(stones [][]int) int {
	var xIndex, yIndex [maxLen][]int
	for i, stone := range stones {
		xIndex[stone[0]] = append(xIndex[stone[0]], i)
		yIndex[stone[1]] = append(yIndex[stone[1]], i)
	}
	var uf = NewUnionFinder(len(stones))
	var visited [maxLen]bool

	var search func(i, root int)
	search = func(i, root int) {
		if visited[i] {
			return
		}
		visited[i] = true
		uf.Union(root, i)
		for _, p := range xIndex[stones[i][0]] {
			search(p, root)
		}
		for _, p := range yIndex[stones[i][1]] {
			search(p, root)
		}
	}
	for i := range stones {
		search(i, i)
	}
	return len(stones) - uf.Count()
}

package p341

type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { panic("not implemented") }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { panic("not implemented") }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) { panic("not implemented") }

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) { panic("not implemented") }

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { panic("not implemented") }

type NestedIterator struct {
	idx            int
	nestedIterator *NestedIterator
	nestedList     []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		idx:        -1,
		nestedList: nestedList,
	}
}

func (ni *NestedIterator) Next() int {
	if nestedIterator := ni.nestedIterator; nestedIterator != nil {
		return nestedIterator.Next()
	}

	return ni.nestedList[ni.idx].GetInteger()
}

func (ni *NestedIterator) HasNext() bool {
	if iterator := ni.nestedIterator; iterator != nil {
		if iterator.HasNext() {
			return true
		}
		ni.nestedIterator = nil
	}

	ni.idx++
	if ni.idx == len(ni.nestedList) {
		return false
	}

	if ni.nestedList[ni.idx].IsInteger() {
		return true
	}
	ni.nestedIterator = Constructor(ni.nestedList[ni.idx].GetList())
	return ni.HasNext()
}

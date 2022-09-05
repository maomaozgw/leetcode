package treenode

type NAtrTree[T comparable] struct {
	Val      T
	Children []*NAtrTree[T]
}

func NewNA[T comparable](items ...*T) *NAtrTree[T] {
	root := &NAtrTree[T]{
		Val:      *items[0],
		Children: []*NAtrTree[T]{},
	}
	var rootChan = make(chan *NAtrTree[T], (len(items)+1)/2)
	rootChan <- root
	i := 1
	var current *NAtrTree[T]
	for {
		if i == len(items) {
			close(rootChan)
			break
		}
		if items[i] == nil {
			current = <-rootChan
			i++
			continue
		}
		node := &NAtrTree[T]{
			Val:      *items[i],
			Children: []*NAtrTree[T]{},
		}
		rootChan <- node
		current.Children = append(current.Children, node)
		i++
	}
	return root
}

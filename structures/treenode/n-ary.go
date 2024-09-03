package treenode

import "strings"

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

func MustNewNAFromLeetCodeStr[T comparable](converter func(string) (T, error), str string) *NAtrTree[T] {
	strItems := strings.Split(str, ",")
	var result = &NAtrTree[T]{}
	var nodeCh = make(chan *NAtrTree[T], 10000)
	nodeCh <- result
	var currentNode = <-nodeCh
	defer close(nodeCh)
	for _, item := range strItems {
		if item == NullStr {
			currentNode = <-nodeCh
			continue
		}
		val, err := converter(item)
		if err != nil {
			panic(err)
		}
		node := &NAtrTree[T]{
			Val: val,
		}
		nodeCh <- node
		currentNode.Children = append(currentNode.Children, node)
	}
	return result.Children[0]
}

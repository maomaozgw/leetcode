package treenode

import (
	"github.com/maomaozgw/leetcode/structures/tools"
)

type BinaryTree[T comparable] struct {
	Val   T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

func NewFullBinaryTree[T comparable](values ...T) *BinaryTree[T] {
	var ptrVals []*T
	for idx := range values {
		ptrVals = append(ptrVals, tools.Ptr[T](values[idx]))
	}
	return NewBinaryTree(ptrVals...)
}

func NewBinaryTreeFromString[T comparable](convert Converter[T], values ...string) *BinaryTree[T] {
	const (
		null = "null"
	)
	if len(values) == 0 {
		return nil
	}
	rootVal := values[0]
	if rootVal == null {
		panic("root node cannot be null")
	}
	rootNode := &BinaryTree[T]{
		Val:   convert(rootVal),
		Left:  nil,
		Right: nil,
	}
	var rootChan = make(chan *BinaryTree[T], (len(values)+1)/2)
	i := 1
	rootChan <- rootNode
	for current := range rootChan {
		if i == len(values) {
			break
		}
		val := values[i]
		if val != null {
			current.Left = &BinaryTree[T]{
				Val: convert(val),
			}
			rootChan <- current.Left
		}
		i++
		if i == len(values) {
			break
		}
		val = values[i]
		if val != null {
			current.Right = &BinaryTree[T]{
				Val: convert(val),
			}
			rootChan <- current.Right
		}
		i++
	}
	close(rootChan)
	return rootNode
}

func NewBinaryTree[T comparable](values ...*T) *BinaryTree[T] {
	if len(values) == 0 {
		return nil
	}
	rootVal := values[0]
	if rootVal == nil {
		panic("root node cannot be nil")
	}
	rootNode := &BinaryTree[T]{
		Val:   *rootVal,
		Left:  nil,
		Right: nil,
	}
	var rootChan = make(chan *BinaryTree[T], (len(values)+1)/2)
	i := 1
	rootChan <- rootNode
	for current := range rootChan {
		if i == len(values) {
			break
		}
		val := values[i]
		if val != nil {
			current.Left = &BinaryTree[T]{
				Val: *val,
			}
			rootChan <- current.Left
		}
		i++
		if i == len(values) {
			break
		}
		val = values[i]
		if val != nil {
			current.Right = &BinaryTree[T]{
				Val: *val,
			}
			rootChan <- current.Right
		}
		i++
	}
	close(rootChan)
	return rootNode
}

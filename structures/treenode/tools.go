package treenode

import (
	"fmt"

	"github.com/disiqueira/gotree"
)

func Equal[T comparable](one *BinaryTree[T], another *BinaryTree[T]) bool {
	if one == nil && another == nil {
		return true
	} else if one == nil || another == nil {
		return false
	}
	if one.Val != another.Val {
		return false
	}
	if left := Equal(one.Left, another.Left); !left {
		return false
	}

	return Equal(one.Right, another.Right)
}

func DeepIter[T comparable](root *BinaryTree[T], callFunc func(node *BinaryTree[T]) bool) bool {
	willContinue := callFunc(root)
	if !willContinue {
		return false
	}
	if root.Left != nil {
		willContinue = DeepIter(root.Left, callFunc)
		if !willContinue {
			return false
		}
	}
	if root.Right != nil {
		willContinue = DeepIter(root.Right, callFunc)
		if !willContinue {
			return false
		}
	}
	return true
}

func String[T comparable](root *BinaryTree[T]) string {

	if root == nil {
		return ""
	}

	node := gotree.New(fmt.Sprintf("%v", root.Val))
	buildTree(node, root)
	return node.Print()
}

func Print[T comparable](root *BinaryTree[T]) {

	if root == nil {
		return
	}

	node := gotree.New(fmt.Sprintf("%v", root.Val))
	buildTree(node, root)
	fmt.Println(node.Print())
}

func buildTree[T comparable](node gotree.Tree, root *BinaryTree[T]) {
	if root.Left != nil {
		leftNode := node.Add(fmt.Sprintf("%v", root.Left.Val))
		buildTree(leftNode, root.Left)
	}
	if root.Right != nil {
		rightNode := node.Add(fmt.Sprintf("%v", root.Right.Val))
		buildTree(rightNode, root.Right)
	}
}

func MustFirstByValue[T comparable](root *BinaryTree[T], val T) *BinaryTree[T] {
	find, ok := FirstByValue(root, val)
	if !ok {
		panic(fmt.Sprintf("cannot find %v", val))
	}
	return find
}

func FirstByValue[T comparable](root *BinaryTree[T], val T) (*BinaryTree[T], bool) {
	if root == nil {
		return nil, false
	}
	if root.Val == val {
		return root, true
	}
	if find, ok := FirstByValue(root.Left, val); ok {
		return find, ok
	}

	if find, ok := FirstByValue(root.Right, val); ok {
		return find, ok
	}
	return nil, false
}

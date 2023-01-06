package p297

import (
	"strconv"
	"strings"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

const (
	null = "null"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var values []string
	if root == nil {
		return ""
	}
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		var newNodes []*TreeNode
		for _, node := range nodes {
			if node == nil {
				values = append(values, null)
				continue
			}
			values = append(values, strconv.Itoa(node.Val))
			newNodes = append(newNodes, node.Left, node.Right)
		}
		nodes = newNodes
	}
	var lastIdx = len(values) - 1
	for ; values[lastIdx] == null; lastIdx-- {

	}
	return strings.Join(values[:lastIdx+1], ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	var values = strings.Split(data, ",")
	rootVal := values[0]
	if rootVal == null {
		panic("root node cannot be null")
	}
	convert := func(s string) int {
		val, _ := strconv.Atoi(s)
		return val
	}
	rootNode := &TreeNode{
		Val:   convert(rootVal),
		Left:  nil,
		Right: nil,
	}
	var rootChan = make(chan *TreeNode, (len(values)+1)/2)
	i := 1
	rootChan <- rootNode
	for current := range rootChan {
		if i == len(values) {
			break
		}
		val := values[i]
		if val != null {
			current.Left = &TreeNode{
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
			current.Right = &TreeNode{
				Val: convert(val),
			}
			rootChan <- current.Right
		}
		i++
	}
	close(rootChan)
	return rootNode
}

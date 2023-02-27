package p427

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func compare(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 != nil && node2 != nil {

	} else {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	if node1.IsLeaf != node2.IsLeaf {
		return false
	}
	return true
}

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Example 1",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[1,0]`),
			},
			want: &Node{
				Val:    false,
				IsLeaf: false,
				TopLeft: &Node{
					Val:    false,
					IsLeaf: true,
				},
				TopRight: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomRight: &Node{
					Val:    false,
					IsLeaf: true,
				},
			},
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]`),
			},
			want: &Node{
				Val:    false,
				IsLeaf: false,
				TopLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				TopRight: &Node{
					Val:    false,
					IsLeaf: false,
					TopLeft: &Node{
						Val:    false,
						IsLeaf: true,
					},
					TopRight: &Node{
						Val:    false,
						IsLeaf: true,
					},
					BottomLeft: &Node{
						Val:    true,
						IsLeaf: true,
					},
					BottomRight: &Node{
						Val:    true,
						IsLeaf: true,
					},
				},
				BottomLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomRight: &Node{
					Val:    false,
					IsLeaf: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct() = %v, want %v", got, tt.want)
			}
		})
	}
}

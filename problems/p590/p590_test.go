package p590

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

func Test_postorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				root: treenode.MustNewNAFromLeetCodeStr[int](strconv.Atoi, `1,null,3,2,4,null,5,6`),
			},
			want: []int{5, 6, 3, 2, 4, 1},
		},
		{
			name: "Example 2",
			args: args{
				root: treenode.MustNewNAFromLeetCodeStr[int](strconv.Atoi, `1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14`),
			},
			want: []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

package p1519

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_countSubTrees(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		labels string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				n:      7,
				edges:  tools.NewGridFromStr(strconv.Atoi, `[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]`),
				labels: "abaedcd",
			},
			want: []int{2, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "Example 2",
			args: args{
				n:      4,
				edges:  tools.NewGridFromStr(strconv.Atoi, `[0,1],[1,2],[0,3]`),
				labels: "bbbb",
			},
			want: []int{4, 2, 1, 1},
		},
		{
			name: "Example 3",
			args: args{
				n:      5,
				edges:  tools.NewGridFromStr(strconv.Atoi, `[0,1],[0,2],[1,3],[0,4]`),
				labels: "aabab",
			},
			want: []int{3, 2, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubTrees(tt.args.n, tt.args.edges, tt.args.labels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSubTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
